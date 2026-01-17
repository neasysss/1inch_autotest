FROM mcr.microsoft.com/playwright:v1.57.0-noble

WORKDIR /app

RUN corepack enable

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

COPY . .

RUN mkdir -p /app/playwright-report /app/test-results

# запуск в виртуальном дисплее
CMD ["bash", "-lc", "xvfb-run --auto-servernum --server-args='-screen 0 1280x720x24' yarn playwright test; yarn playwright show-report --host 0.0.0.0 --port 5225 playwright-report"]
