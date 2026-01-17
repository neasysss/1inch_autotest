SHELL := /bin/bash

REPORT_PORT := 5225
IMAGE := 1inch-autotest
ENV_FILE := ./.env

.PHONY: local local-ui docker

local:
	yarn playwright test

local-ui:
	yarn playwright test --ui

# Docker: собираем образ, запускаем контейнер в foreground:
# - прогоняет тесты
# - затем поднимает show-report внутри контейнера на http://localhost:5225
# Остановить: Ctrl+C
docker:
	@mkdir -p playwright-report test-results
	docker build -t $(IMAGE) .
	docker run --rm -it --init --ipc=host \
		-p $(REPORT_PORT):$(REPORT_PORT) \
		-e IN_DOCKER=true \
		--env-file $(ENV_FILE) \
		-v "$$(pwd)/playwright-report:/app/playwright-report" \
		-v "$$(pwd)/test-results:/app/test-results" \
		$(IMAGE)
