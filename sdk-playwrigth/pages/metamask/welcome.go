package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type Welcome struct {
	page playwright.Page
}

func NewWelcome(page playwright.Page) *Welcome {
	return &Welcome{
		page: page,
	}
}

func (w *Welcome) HaveExistingWalletClick() error {
	err := w.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	return nil
}

func (w *Welcome) ImportUsingPhraseClick() (*ImportWithPhrase, error) {
	err := w.page.Locator(`//html/body/div[3]/div[2]/div/section/div/button[3]`).Click()
	if err != nil {
		return nil, fmt.Errorf("клик кнопки: %w", err)
	}

	page := NewImportWithPhrase(w.page)

	return page, nil
}
