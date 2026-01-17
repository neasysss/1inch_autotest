import { defineConfig, devices } from "@playwright/test";
import dotenv from "dotenv-safe";

dotenv.config();

/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
  testDir: "./tests",
  fullyParallel: true,
  forbidOnly: !!process.env.IN_DOCKER,
  retries: 0,
  workers: 1,
  /* Reporter to use. See https://playwright.dev/docs/test-reporters */
  reporter: process.env.IN_DOCKER
  ? [["list", { printSteps: true }], ["html", { open: "never" }]]
  : [["html", { open: 'always' }], ["list", { printSteps: true }]],
  /* Shared settings for all the projects below. See https://playwright.dev/docs/api/class-testoptions. */
  use: {
    baseURL: process.env.FRONTEND_URL,
    trace: "on-first-retry",
    headless: false,
    screenshot: {
      mode: "on",
      fullPage: true,
    },
  },
});
