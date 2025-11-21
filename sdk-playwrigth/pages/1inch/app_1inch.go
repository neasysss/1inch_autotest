package oneInch

import (
	"fmt"
	"slices"
	"time"

	"github.com/playwright-community/playwright-go"
)

type App1inch struct {
	page playwright.Page
}

func NewApp1inch(page playwright.Page) *App1inch {
	return &App1inch{
		page: page,
	}
}

func (a *App1inch) ConnectMetamask() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/app-header/div[1]/div/div[3]/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	err = a.page.Locator(`//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-wallet-list/div[1]/oi-wallet-cell[3]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	err = a.page.Locator(`//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-select-network/div/div[1]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	return nil
}

func (a *App1inch) OpenToken1Selector() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field/oi-amount-input/div[1]/button`).Click()
	if err != nil {
		return fmt.Errorf("клик по выбору первого токена: %w", err)
	}

	return nil
}

func (a *App1inch) SelectArbitrum() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-chain-selector/button`).Click()
	if err != nil {
		return fmt.Errorf("клик по выбору сети: %w", err)
	}

	err = a.page.Locator(`//*[@id="oi-select-item-11"]`).Click()
	if err != nil {
		return fmt.Errorf("выбор арбитрума: %w", err)
	}

	return nil
}

func (a *App1inch) SelectTokenOne(token string) error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div/oi-search/div/oi-input/label/input`).Fill(token)
	if err != nil {
		return fmt.Errorf("ввод первого токена: %w", err)
	}

	time.Sleep(time.Second * 3)

	token1ListLocator := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[2]/oi-tokens-list/div`)

	token1Locator, err := getLocatorWithDivText(token1ListLocator, token)
	if err != nil {
		return fmt.Errorf("поиск первого токена: %w", err)
	}

	err = token1Locator.Click()
	if err != nil {
		return fmt.Errorf("выбор первого токена: %w", err)
	}

	return nil
}

func (a *App1inch) OpenToken2Selector() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-select-token/div[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик по выбору второго токена: %w", err)
	}

	return nil
}

func (a *App1inch) SelectTokenTwo(token string) error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-search/div/oi-input/label`).Fill(token)
	if err != nil {
		return fmt.Errorf("ввода второго токена: %w", err)
	}

	time.Sleep(time.Second * 3)

	token2ListLocator := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[2]/oi-tokens-list/div`)

	token2Locator, err := getLocatorWithDivText(token2ListLocator, token)
	if err != nil {
		return fmt.Errorf("поиск второго токена: %w", err)
	}

	err = token2Locator.Click()
	if err != nil {
		return fmt.Errorf("выбор второго токена %s: %w", token, err)
	}

	return nil
}

func (a *App1inch) InputAmount(sum string) error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field[1]/oi-amount-input/div[1]/input`).Fill(sum)
	if err != nil {
		return fmt.Errorf("ввод суммы: %w", err)
	}

	return nil
}

func (a *App1inch) MaxSum() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field[1]/div/div[2]/span[1]`).Click()
	if err != nil {
		return fmt.Errorf("клик максимальной суммы: %w", err)
	}

	return nil
}

func (a *App1inch) Swap() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/div[2]/oi-swap-button/button`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	return nil
}

func (a *App1inch) FlipTokens() error {
	err := a.page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/div[1]/button`).Click()
	if err != nil {
		return fmt.Errorf("клик флипа токенов: %w", err)
	}

	return nil
}

func getLocatorWithDivText(source playwright.Locator, text string) (playwright.Locator, error) {
	listFiltered, err := source.Locator("div").Filter(playwright.LocatorFilterOptions{
		HasText: text,
	}).All()
	if err != nil {
		return nil, fmt.Errorf("поиск объектов: %w", err)
	}

	var result playwright.Locator

	for _, elem := range listFiltered {
		textContents, err := elem.AllTextContents()
		if err != nil {
			return nil, fmt.Errorf("получение текста: %w", err)
		}

		if slices.Contains(textContents, text) {
			result = elem
			break
		}
	}

	if result == nil {
		return nil, fmt.Errorf("div с текстом %s не найден", text)
	}

	return result, nil
}
