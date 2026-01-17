import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";

export class OneInchSwapSumbitted extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
  }

  private SELECTORS = {
    SWAP_SUMBITTED_CLOSE_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-information-dialog/div[2]/button')  
  }

  async swapSubmittedCloseClick() {
    await this.SELECTORS.SWAP_SUMBITTED_CLOSE_BUTTON.click()
  }

  async verifyRequiredElementsPresent(): Promise<void> {
    await expect(this.SELECTORS.SWAP_SUMBITTED_CLOSE_BUTTON).toBeVisible({
        timeout: 10 * 1000 // 10 секунд
    })
  }
}
