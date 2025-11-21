package browser

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func LaunchChromium() (playwright.BrowserContext, func(), error) {
	pw, err := playwright.Run(&playwright.RunOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("запуск plywright: %w", err)
	}

	browser, err := pw.Chromium.LaunchPersistentContext("", playwright.BrowserTypeLaunchPersistentContextOptions{
		Headless: playwright.Bool(false),
		Args: []string{
			"--disable-extensions-except=./metamask",
			"--load-extension=./metamask",
			"--disable-web-security",
			"--no-sandbox",
			"--allow-file-access-from-files",
		},
	})
	if err != nil {
		pw.Stop()

		return nil, nil, fmt.Errorf("запуск chromium: %w", err)
	}

	// 15 сек
	browser.SetDefaultTimeout(15000)

	stop := func() {
		pw.Stop()
		browser.Close()
	}

	return browser, stop, nil
}

func GetLastPage(browser playwright.BrowserContext) playwright.Page {
	return browser.Pages()[len(browser.Pages())-1]
}
