import dotenv from "dotenv-safe";
import { createMetamaskFixture } from "../../../fixtures/metaMaskFixtures";

import { OneInchSwap } from "../pages/1inch/Swap";
import { OneInchSwapSumbitted} from "../pages/1inch/SwapSubmitted"
import { MetamaskWelcome } from "../pages/metamask/Welcome";
import { MetamaskConnectConfirm } from "../pages/metamask/ConnectConfirm";
import { MetamaskAcceptNetwork } from "../pages/metamask/AcceptNetwork";
import { MetamaskSwapConfirm } from "../pages/metamask/SwapConfirm";

dotenv.config();

const {test: test1inch} = createMetamaskFixture()

test1inch.describe("1inch swap and back", () => {
  test1inch.describe.configure({ mode: "serial" });

  test1inch("import metamask", async ({context, seedPhrase, password}) => {
    await new Promise(resolve => setTimeout(resolve, 5 * 1000)); // ждем пока откроется метамаск

    var mmWelcomePage = new MetamaskWelcome(context.pages()[1]) 

    await mmWelcomePage.verifyRequiredElementsPresent()
    await mmWelcomePage.haveExistingWalletClick();

    var mmImportWithPhrasePage = await mmWelcomePage.importUsingPhraseClick();
    
    await mmImportWithPhrasePage.seedEnter(seedPhrase)
    
    var mmCreatePasswordPage = await mmImportWithPhrasePage.continueClick();
    
    await mmCreatePasswordPage.passwordEnter(password)
    await mmCreatePasswordPage.termsAgreeClick();
    
    var mmMetricsPage = await mmCreatePasswordPage.createPasswordClick();
    
    await mmMetricsPage.usageDataCheckboxClick()
    
    var mmCompletionPage = await mmMetricsPage.continueClick()
    
    await mmCompletionPage.doneClick()
  })

  test1inch("connect wallet", async ({context}) => {
    var inchSwapPage = new OneInchSwap(await context.newPage())

    await inchSwapPage.navigateTo()
    await inchSwapPage.verifyRequiredElementsPresent()
    await inchSwapPage.connectMetamask()
    await new Promise(resolve => setTimeout(resolve, 3 * 1000)); // ждем открытия вкладки подключения кошелька
    
    var mmConnectConfirmPage = new MetamaskConnectConfirm(context.pages()[3]) 
    
    await mmConnectConfirmPage.verifyRequiredElementsPresent()
    await mmConnectConfirmPage.connectConfirmClick()
  })

  test1inch("configure tokens",  async ({context, network, token1, token2, amount}) => {
    var inchSwapPage = new OneInchSwap(context.pages()[2])
  
    await inchSwapPage.openToken1Selector()
    await inchSwapPage.selectNetwork(network)
    await inchSwapPage.selectToken(token1)

    await inchSwapPage.openToken2Selector()
    await inchSwapPage.selectToken(token2)

    await inchSwapPage.inputAmount(amount)
  })

  test1inch("do swap", async ({context}) => {
    var inchSwapPage = new OneInchSwap(context.pages()[2])
    
    await inchSwapPage.verifyRequiredElementsPresent()
    await inchSwapPage.swapClick()
    await new Promise(resolve => setTimeout(resolve, 3 * 1000)); // ждем открытия вкладки подтверждения сети
    
    var mmAcceptNetworkPage = new MetamaskAcceptNetwork(context.pages()[3]) 

    await mmAcceptNetworkPage.verifyRequiredElementsPresent()
    await mmAcceptNetworkPage.confirmClick()
    await new Promise(resolve => setTimeout(resolve, 7 * 1000)); // ждем открытия вкладки подтверждения транзакции

    var mmSwapConfirmPage = new MetamaskSwapConfirm(context.pages()[3])

    await mmSwapConfirmPage.verifyRequiredElementsPresent()
    await mmSwapConfirmPage.confirmClick()
  })

  test1inch("swap wait", async ({context}) => {
    await new Promise(resolve => setTimeout(resolve, 5 * 1000)); // ждем своп

    var inchSwapSubmittedPage = new OneInchSwapSumbitted(context.pages()[2])

    await inchSwapSubmittedPage.verifyRequiredElementsPresent()
    await inchSwapSubmittedPage.swapSubmittedCloseClick()
  })

  test1inch("do swap back", async ({context}) => {
    var inchSwapPage = new OneInchSwap(context.pages()[2])

    await inchSwapPage.verifyRequiredElementsPresent()
    await inchSwapPage.flipTokensClick()
    await inchSwapPage.inputAmount('max')
    await inchSwapPage.swapClick()
    await new Promise(resolve => setTimeout(resolve, 7 * 1000)); // ждем открытия вкладки подтверждения транзакции

    var mmSwapConfirmPage = new MetamaskSwapConfirm(context.pages()[3])

    await mmSwapConfirmPage.verifyRequiredElementsPresent()
    await mmSwapConfirmPage.confirmClick()
  })

  test1inch("swap back wait", async ({context}) => {
    await new Promise(resolve => setTimeout(resolve, 5 * 1000)); // ждем своп

    var inchSwapSubmittedPage = new OneInchSwapSumbitted(context.pages()[2])

    await inchSwapSubmittedPage.verifyRequiredElementsPresent()
    await inchSwapSubmittedPage.swapSubmittedCloseClick()
  })
});
