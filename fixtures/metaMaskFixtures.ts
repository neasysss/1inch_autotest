import { test as base, chromium, type BrowserContext } from "@playwright/test";
import path from "path";

type TestFixtures = {
  context: BrowserContext;
  seedPhrase: string
  password: string
  token1: string
  token2: string
  amount: string
};

type WorkerFixtures = {
  sharedContext: BrowserContext;
};

export function createMetamaskFixture() {
  const test = base.extend<TestFixtures, WorkerFixtures>({
    sharedContext: [
      async ({}, use) => {
        const pathToExtension = path.join(
          __dirname,
          "../testData/extensions/metamask"
        );

        const ctx = await chromium.launchPersistentContext("", {
          headless: false,
          args: [
            `--disable-extensions-except=${pathToExtension}`,
            `--load-extension=${pathToExtension}`,
          ],
        });

        await use(ctx);
        await ctx.close();
      },
      { scope: "worker" },
    ],
    
    context: async ({ sharedContext }, use) => {
      await use(sharedContext);
    },
    seedPhrase: async ({}, use) => {      
      use(process.env.METAMASK_SEED)
    }, 
    password: async ({}, use) => {
      use(process.env.METAMASK_PASSWORD)
    }, 
    token1: async ({}, use) => {
      use(process.env.TOKEN_1)
    }, 
    token2: async ({}, use) => {
      use(process.env.TOKEN_2)
    },
    amount: async ({}, use) => {
      use(process.env.AMOUNT)
    },
  });

  return { test, expect: test.expect };
}