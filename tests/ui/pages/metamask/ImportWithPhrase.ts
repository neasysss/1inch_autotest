import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";
import { MetamaskCreatePassword } from "./CreatePasword";

export class MetamaskImportWithPhrase extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "chrome-extension://idngmeffehjdaahfoiejjnpahdedomlh/home.html#onboarding/import-with-recovery-phrase";
  }

  private SELECTORS = {
    PHRASE_ENTER_FIELD: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[1]/div[4]/form/div/div[1]/div/textarea'),
    CONTINUE_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button')
  }

  async seedEnter(seed: string) {
    await this.SELECTORS.PHRASE_ENTER_FIELD.pressSequentially(seed);
  }

  async continueClick()  {
    await this.SELECTORS.CONTINUE_BUTTON.click();

    var nextPage = new MetamaskCreatePassword(this.page)
    await nextPage.verifyRequiredElementsPresent()

    return nextPage
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.PHRASE_ENTER_FIELD).toBeVisible({timeout:timeout})
    await expect(this.SELECTORS.CONTINUE_BUTTON).toBeVisible({timeout:timeout})
  }
}
