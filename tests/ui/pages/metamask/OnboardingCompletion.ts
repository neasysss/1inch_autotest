import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";

export class MetamaskOnboardingCompletion extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "chrome-extension://idngmeffehjdaahfoiejjnpahdedomlh/home.html#onboarding/completion";
  }

  private SELECTORS = {
    DONE_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button')
  }

  async doneClick() {
    await this.SELECTORS.DONE_BUTTON.click();
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.DONE_BUTTON).toBeVisible({timeout:timeout})
  }
}
