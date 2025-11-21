install_pw:
	go run github.com/playwright-community/playwright-go/cmd/playwright@v0.5200.1 install --with-deps

e2e: install_pw
	go test -v ./...