import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";

export class MetamaskConnectConfirm extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
  }

  private SELECTORS = {
    CONNECT_CONFIRM_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/div/div[3]/div/div/button[2]'),
  }

  async connectConfirmClick() {
    await this.SELECTORS.CONNECT_CONFIRM_BUTTON.click();
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.CONNECT_CONFIRM_BUTTON).toBeVisible({timeout:timeout})
  }
}
