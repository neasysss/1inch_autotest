package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type CreatePassword struct {
	page playwright.Page
}

func NewCreatePassword(page playwright.Page) *CreatePassword {
	return &CreatePassword{
		page: page,
	}
}

func (c *CreatePassword) PasswordFill(password string) error {
	err := c.page.Locator(`//*[@id="create-password-new"]`).Fill(password)
	if err != nil {
		return fmt.Errorf("ввод пароля: %w", err)
	}

	err = c.page.Locator(`//*[@id="create-password-confirm"]`).Fill(password)
	if err != nil {
		return fmt.Errorf("ввод подтверждения пароля: %w", err)
	}

	return nil
}

func (c *CreatePassword) AgreeTerms() error {
	err := c.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/form/div[1]/div[4]/label/span[1]/input`).Click()
	if err != nil {
		return fmt.Errorf("клик чекбокса: %w", err)
	}

	return nil
}

func (c *CreatePassword) CreatePassword() (*Metametrics, error) {
	err := c.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/form/div[2]/button`).Click()
	if err != nil {
		return nil, fmt.Errorf("клик кнопки: %w", err)
	}

	page := NewMetametrics(c.page)

	return page, nil
}
