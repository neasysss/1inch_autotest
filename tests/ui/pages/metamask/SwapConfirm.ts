import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";

export class MetamaskSwapConfirm extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
  }

  private SELECTORS = {
    CONFIRM_BUTTON: this.page.locator('//*[@id="app-content"]/div/div[2]/div/div/div[3]/div/button[2]')
  }

  async confirmClick() {
    await this.SELECTORS.CONFIRM_BUTTON.click();
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.CONFIRM_BUTTON).toBeVisible({timeout:timeout})
  }
}
