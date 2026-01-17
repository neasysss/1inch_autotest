import { expect, Page } from "@playwright/test";
import { BasePage } from "../BasePage";

export class OneInchSwap extends BasePage {
  constructor(page: Page) {
    super(page);
    this.page = page;
    this.url = "https://1inch.com/swap?src=1:ETH";
  }

  private SELECTORS = {
    SWAP_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/div[2]/oi-swap-button/button'),
    CONNECT_WALLET_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/app-header/div[1]/div/div[3]/button[2]'),
    CONNECT_WALLET_VIEW: {
      WALLETS_LIST: this.page.locator('//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-wallet-list/div[1]'),
      EVM_BASED_BUTTON: this.page.locator('//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-select-network/div/div[1]')
    },
    TOKEN1_SELECT_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field/oi-amount-input/div[1]/button'),
    TOKEN2_SELECT_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-select-token/div[2]'),
    TOKEN_SELECTOR_VIEW: {
      NETWORK_SELECTOR_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-chain-selector/button'),
      NETWORK_SELECTOR_VIEW: {
        NETWORKS_LIST: this.page.locator('.cdk-overlay-pane:visible oi-select .select-viewport')
      },
      SEARCH_FIELD: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-search/div/oi-input/label/input'),
      SEARCH_RESULT_LIST: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[2]/oi-tokens-list[1]/div')
    },
    AMOUNT_FIELD: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field[1]/oi-amount-input/div[1]/input'),
    MAX_AMOUNT_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field[1]/div/div[2]/span[1]'),
    FLIP_TOKENS_BUTTON: this.page.locator('//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/div[1]/button')
  }

  async connectMetamask() {
    await this.SELECTORS.CONNECT_WALLET_BUTTON.click()
    await this.SELECTORS.CONNECT_WALLET_VIEW.WALLETS_LIST.getByText('MetaMask', {exact: true}).click()
    await this.SELECTORS.CONNECT_WALLET_VIEW.EVM_BASED_BUTTON.click()
  }

  async openToken1Selector() {
    await this.SELECTORS.TOKEN1_SELECT_BUTTON.click()
  }

  async openToken2Selector() {
    await this.SELECTORS.TOKEN2_SELECT_BUTTON.click()
  }

  async selectNetwork(network: string) {
    await this.SELECTORS.TOKEN_SELECTOR_VIEW.NETWORK_SELECTOR_BUTTON.click()
    await this.SELECTORS.TOKEN_SELECTOR_VIEW.NETWORK_SELECTOR_VIEW.NETWORKS_LIST.
      getByText(network, {exact: true}).click()
  }

  async selectToken(token: string) {
    await this.SELECTORS.TOKEN_SELECTOR_VIEW.SEARCH_FIELD.fill(token)
    await this.SELECTORS.TOKEN_SELECTOR_VIEW.SEARCH_RESULT_LIST.getByText(token, {exact: true}).click()
  }

  async inputAmount(amount: string) {
    if (amount == 'max') {
      await this.SELECTORS.AMOUNT_FIELD.hover() // наводим мышь на поле ввода, чтобы появилась кнопка max
      await new Promise(resolve => setTimeout(resolve, 1 * 1000)); // ждем пока инч корректно подгрузит баланс
      await this.SELECTORS.MAX_AMOUNT_BUTTON.click()
    } else {
      await this.SELECTORS.AMOUNT_FIELD.fill(amount)
    }
  }

  async flipTokensClick() {
    await this.SELECTORS.FLIP_TOKENS_BUTTON.click()
  }

  async swapClick() {
    await this.SELECTORS.SWAP_BUTTON.click()
  }

  async navigateTo(): Promise<void> {
    await this.page.goto(this.url);
  }

  async verifyRequiredElementsPresent(): Promise<void> {
    await expect(this.SELECTORS.SWAP_BUTTON).toBeVisible({
        timeout: 10 * 1000 // 10 секунд
    })
  }
}
