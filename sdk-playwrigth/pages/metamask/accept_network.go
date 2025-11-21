package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type AcceptNetwork struct {
	page playwright.Page
}

func NewAcceptNetwork(page playwright.Page) *AcceptNetwork {
	return &AcceptNetwork{
		page: page,
	}
}

func (a *AcceptNetwork) Confirm() error {
	//одобрить
	err := a.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div[3]/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	return nil
}
