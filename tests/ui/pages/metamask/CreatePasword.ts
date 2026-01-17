import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";
import { MetamaskMetrics } from "./MetaMetrics";

export class MetamaskCreatePassword extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "chrome-extension://idngmeffehjdaahfoiejjnpahdedomlh/home.html#onboarding/create-password";
  }

  private SELECTORS = {
    NEW_PASSWORD_FIELD: this.page.locator('//*[@id="create-password-new"]'),
    CONFIRM_PASSWORD_FIELD: this.page.locator('//*[@id="create-password-confirm"]'),
    TERMS_CHECKBOX: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/form/div[1]/div[4]/label/span[1]/input'),
    CREATE_PASSWORD_BUTTON: this.page.locator('//*[@id="app-content"]/div/div/div/div[2]/form/div[2]/button')
  }

  async passwordEnter(password: string) {
    await this.SELECTORS.NEW_PASSWORD_FIELD.fill(password)
    await this.SELECTORS.CONFIRM_PASSWORD_FIELD.fill(password)
  }

  async termsAgreeClick() {
    await this.SELECTORS.TERMS_CHECKBOX.click();
  }

  async createPasswordClick()  {
    await this.SELECTORS.CREATE_PASSWORD_BUTTON.click();

    var nextPage = new MetamaskMetrics(this.page)
    await nextPage.verifyRequiredElementsPresent()

    return nextPage
  }

  async navigateTo() {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent() {
    var timeout = 10*1000
    await expect(this.SELECTORS.NEW_PASSWORD_FIELD).toBeVisible({timeout:timeout})
    await expect(this.SELECTORS.CONFIRM_PASSWORD_FIELD).toBeVisible({timeout:timeout})
    await expect(this.SELECTORS.TERMS_CHECKBOX).toBeVisible({timeout:timeout})
    await expect(this.SELECTORS.CREATE_PASSWORD_BUTTON).toBeVisible({timeout:timeout})
  }
}
