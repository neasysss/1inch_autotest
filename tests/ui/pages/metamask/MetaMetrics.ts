import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";
import { MetamaskOnboardingCompletion } from "./OnboardingCompletion";

export class MetamaskMetrics extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "chrome-extension://idngmeffehjdaahfoiejjnpahdedomlh/home.html#onboarding/metametrics";
  }

  private SELECTORS = {
    USAGE_DATA_CHECKBOX: this.page.locator('//*[@id="metametrics-opt-in"]'),
    CONTINUE_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[4]/button')
  }

  async usageDataCheckboxClick() {
    await this.SELECTORS.USAGE_DATA_CHECKBOX.click();
  }

  async continueClick()  {
    await this.SELECTORS.CONTINUE_BUTTON.click();

    var nextPage = new MetamaskOnboardingCompletion(this.page)
    await nextPage.verifyRequiredElementsPresent()

    return nextPage
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.USAGE_DATA_CHECKBOX).toBeVisible({timeout:timeout})
    await expect(this.SELECTORS.CONTINUE_BUTTON).toBeVisible({timeout:timeout})
  }
}
