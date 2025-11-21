package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type SwapConfirm struct {
	page playwright.Page
}

func NewSwapConfirm(page playwright.Page) *SwapConfirm {
	return &SwapConfirm{
		page: page,
	}
}

func (s *SwapConfirm) Confirm() error {
	err := s.page.Locator(`//*[@id="app-content"]/div/div[2]/div/div/div[3]/div/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки: %w", err)
	}
	return nil
}
