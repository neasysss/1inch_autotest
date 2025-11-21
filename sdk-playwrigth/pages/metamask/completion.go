package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type Completion struct {
	page playwright.Page
}

func NewCompletion(page playwright.Page) *Completion {
	return &Completion{
		page: page,
	}
}

func (c *Completion) Done() error {
	err := c.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}

	return nil
}
