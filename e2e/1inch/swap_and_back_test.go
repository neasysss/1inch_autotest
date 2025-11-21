package oneInch

import (
	"1inch_autotest/sdk-playwrigth/browser"
	oneInch "1inch_autotest/sdk-playwrigth/pages/1inch"
	"1inch_autotest/sdk-playwrigth/pages/metamask"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestSwapAndBack(t *testing.T) {
	args, err := loadArgs()
	require.NoError(t, err)

	chromium, stopFn, err := browser.LaunchChromium()
	require.NoError(t, err)
	t.Cleanup(stopFn)

	t.Log("браузер запущен")

	// ожидание открытия страницы с расширением
	time.Sleep(time.Second * 3)

	require.True(t, t.Run("import_metamask", func(t *testing.T) {
		var metamaskPage playwright.Page

		pages := chromium.Pages()
		for _, page := range pages {
			if strings.HasPrefix(page.URL(), "chrome-extension://") {
				metamaskPage = page
			}
		}
		require.NotEmpty(t, metamaskPage)

		err = importMetamask(metamaskPage, args.seed, args.password)
		require.NoError(t, err)
	}))

	var oneInchPage playwright.Page

	require.True(t, t.Run("open_1inch", func(t *testing.T) {
		oneInchPage, err = chromium.NewPage()
		require.NoError(t, err)

		_, err = oneInchPage.Goto("https://1inch.com/swap")
		require.NoError(t, err)

		err = oneInchPage.WaitForLoadState()
		require.NoError(t, err)
	}))

	require.True(t, t.Run("connect_wallet", func(t *testing.T) {
		err = connectWallet(oneInchPage)
		require.NoError(t, err)

		t.Log("выполняется запрос на подключение кошелька")
	}))

	// ждем пока браузер откроет окно с подтверждением
	time.Sleep(time.Second * 3)

	require.True(t, t.Run("accept_connect_wallet", func(t *testing.T) {
		err = acceptConnectWallet(browser.GetLastPage(chromium))
		require.NoError(t, err)

		t.Log("кошелек подключен")
	}))

	require.True(t, t.Run("select_tokens", func(t *testing.T) {
		err = selectTokens(oneInchPage, args.token1, args.token2)
		require.NoError(t, err)

		t.Log("токены для свопа выбраны")
	}))

	require.True(t, t.Run("swap", func(t *testing.T) {
		err = swapToken(oneInchPage, args.amount)
		require.NoError(t, err)

		t.Log("сумма для обмена введена")

		time.Sleep(3 * time.Second)

		err = acceptNetwork(browser.GetLastPage(chromium))
		require.NoError(t, err)

		// ожидаем открытия подтверждения свопа
		time.Sleep(5 * time.Second)

		err = acceptSwap(browser.GetLastPage(chromium))
		require.NoError(t, err)
	}))

	// ожидаем своп
	time.Sleep(5 * time.Second)

	require.True(t, t.Run("swap_back", func(t *testing.T) {
		err = flipTokens(oneInchPage)
		require.NoError(t, err)

		// ожидаем появления подсказок по сумме
		time.Sleep(3 * time.Second)

		err = swapToken(oneInchPage, "max")
		require.NoError(t, err)

		// ожидаем открытия подтверждения свопа
		time.Sleep(5 * time.Second)

		err = acceptSwap(browser.GetLastPage(chromium))
		require.NoError(t, err)
	}))
}

func importMetamask(page playwright.Page, seed string, password string) error {
	welcomePage := metamask.NewWelcome(page)

	err := welcomePage.HaveExistingWalletClick()
	if err != nil {
		return fmt.Errorf("выбор существующего кошелька: %w", err)
	}

	importPage, err := welcomePage.ImportUsingPhraseClick()
	if err != nil {
		return fmt.Errorf("импорт с помощью seed-фразы: %w", err)
	}

	err = importPage.EnterSecretPhrase(seed)
	if err != nil {
		return fmt.Errorf("ввод seed-фразы: %w", err)
	}

	createPasswordPage, err := importPage.Continue()
	if err != nil {
		return fmt.Errorf("переход далее: %w", err)
	}

	err = createPasswordPage.PasswordFill(password)
	if err != nil {
		return fmt.Errorf("ввод пароля: %w", err)
	}

	err = createPasswordPage.AgreeTerms()
	if err != nil {
		return fmt.Errorf("соглашение с условиями: %w", err)
	}

	metametricsPage, err := createPasswordPage.CreatePassword()
	if err != nil {
		return fmt.Errorf("создание пароля: %w", err)
	}

	completionPage, err := metametricsPage.Disagree()
	if err != nil {
		return fmt.Errorf("отказ от метрик: %w", err)
	}

	err = completionPage.Done()
	if err != nil {
		return fmt.Errorf("подтверждение: %w", err)
	}

	return nil
}

func connectWallet(page playwright.Page) error {
	oneInchPage := oneInch.NewApp1inch(page)

	err := oneInchPage.ConnectMetamask()
	if err != nil {
		return fmt.Errorf("подключение metamask: %w", err)
	}

	return nil
}

func acceptConnectWallet(page playwright.Page) error {
	connectConfirmPage := metamask.NewConnectConfirm(page)

	err := connectConfirmPage.AcceptConnectWallet()
	if err != nil {
		return fmt.Errorf("подключение кошелька: %w", err)
	}
	return nil
}

func acceptNetwork(page playwright.Page) error {
	acceptPage := metamask.NewAcceptNetwork(page)

	err := acceptPage.Confirm()
	if err != nil {
		return fmt.Errorf("подтверждение сети: %w", err)
	}

	return nil
}

func acceptSwap(page playwright.Page) error {
	acceptSwapTokenTwo := metamask.NewSwapConfirm(page)

	err := acceptSwapTokenTwo.Confirm()
	if err != nil {
		return fmt.Errorf("подтверждение свопа: %w", err)
	}

	return nil
}

func flipTokens(page playwright.Page) error {
	oneInchPage := oneInch.NewApp1inch(page)

	err := oneInchPage.FlipTokens()
	if err != nil {
		return fmt.Errorf("флип токенов: %w", err)
	}
	return nil
}

func selectTokens(page playwright.Page, token1 string, token2 string) error {
	oneInchPage := oneInch.NewApp1inch(page)

	err := oneInchPage.OpenToken1Selector()
	if err != nil {
		return fmt.Errorf("открытие селектора первого токена: %w", err)
	}

	err = oneInchPage.SelectArbitrum()
	if err != nil {
		return fmt.Errorf("выбор арбитрума: %w", err)
	}

	err = oneInchPage.SelectTokenOne(token1)
	if err != nil {
		return fmt.Errorf("выбор первого токена: %w", err)
	}

	err = oneInchPage.OpenToken2Selector()
	if err != nil {
		return fmt.Errorf("открытие селектора второго токена: %w", err)
	}

	err = oneInchPage.SelectTokenTwo(token2)
	if err != nil {
		return fmt.Errorf("выбор второго токена: %w", err)
	}

	return nil
}

func swapToken(page playwright.Page, sum string) error {
	oneInchPage := oneInch.NewApp1inch(page)

	if sum == "max" {
		err := oneInchPage.MaxSum()
		if err != nil {
			return fmt.Errorf("выбор максимальной суммы: %w", err)
		}
	} else {
		err := oneInchPage.InputAmount(sum)
		if err != nil {
			return fmt.Errorf("ввод суммы: %w", err)
		}
	}

	err := oneInchPage.Swap()
	if err != nil {
		return fmt.Errorf("клик свопа: %w", err)

	}

	return nil
}
