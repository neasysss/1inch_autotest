package metamask

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type Metametrics struct {
	page playwright.Page
}

func NewMetametrics(page playwright.Page) *Metametrics {
	return &Metametrics{
		page: page,
	}
}

func (m *Metametrics) Disagree() (*Completion, error) {
	err := m.page.Locator(`//*[@id="metametrics-opt-in"]`).Click()
	if err != nil {
		return nil, fmt.Errorf("клик чекбокса: %w", err)
	}

	err = m.page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[4]/button`).Click()
	if err != nil {
		return nil, fmt.Errorf("клик кнопки: %w", err)
	}

	page := NewCompletion(m.page)

	return page, nil
}
