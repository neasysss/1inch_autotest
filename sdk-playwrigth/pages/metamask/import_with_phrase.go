package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type ImportWithPhrase struct {
	page playwright.Page
}

func NewImportWithPhrase(page playwright.Page) *ImportWithPhrase {
	return &ImportWithPhrase{
		page: page,
	}
}

func (i *ImportWithPhrase) EnterSecretPhrase(seed string) error {
	err := i.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[1]/div[4]/form/div/div[1]/div/textarea`).PressSequentially(seed)
	if err != nil {
		return fmt.Errorf("ввод seed-фразы: %w", err)
	}

	return nil
}

func (i *ImportWithPhrase) Continue() (*CreatePassword, error) {
	err := i.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button`).Click()
	if err != nil {
		return nil, fmt.Errorf("клик кнопки: %w", err)
	}
	page := NewCreatePassword(i.page)

	return page, nil
}
