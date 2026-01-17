import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";
import { MetamaskImportWithPhrase } from "./ImportWithPhrase";

export class MetamaskWelcome extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "chrome-extension://idngmeffehjdaahfoiejjnpahdedomlh/home.html";
  }

  private SELECTORS = {
    HAVE_EXISTING_WALLET_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button[2]'),
    IMPORT_USING_PHRASE_BUTTON: this.page.locator('//html/body/div[3]/div[2]/div/section/div/button[3]')
  }

  async haveExistingWalletClick() {
    await this.SELECTORS.HAVE_EXISTING_WALLET_BUTTON.click();
  }

  async importUsingPhraseClick() {
    await this.SELECTORS.IMPORT_USING_PHRASE_BUTTON.click();

    var nextPage = new MetamaskImportWithPhrase(this.page)
    await nextPage.verifyRequiredElementsPresent()

    return nextPage
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.HAVE_EXISTING_WALLET_BUTTON).toBeVisible({timeout:timeout})
  }
}
