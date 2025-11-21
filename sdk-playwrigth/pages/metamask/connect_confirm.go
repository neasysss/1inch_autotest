package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type ConnectConfirm struct {
	page playwright.Page
}

func NewConnectConfirm(page playwright.Page) *ConnectConfirm {
	return &ConnectConfirm{
		page: page,
	}
}

func (c *ConnectConfirm) AcceptConnectWallet() error {
	err := c.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[3]/div/div/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("клик кнопки:: %w", err)
	}
	return nil
}
