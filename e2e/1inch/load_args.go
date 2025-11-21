package oneInch

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type testArgs struct {
	seed     string
	token1   string
	token2   string
	amount   string
	password string
}

func loadArgs() (testArgs, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return testArgs{}, fmt.Errorf("загрузка .env файла: %w", err)
	}

	seed := os.Getenv("METAMASK_SEED")
	if seed == "" {
		return testArgs{}, fmt.Errorf("отсутствует METAMASK_SEED")
	}

	token1 := os.Getenv("TOKEN_1")
	if token1 == "" {
		return testArgs{}, fmt.Errorf("отсутствует TOKEN_1")
	}

	token2 := os.Getenv("TOKEN_2")
	if token2 == "" {
		return testArgs{}, fmt.Errorf("отсутствует TOKEN_2")
	}

	amount := os.Getenv("AMOUNT")
	if amount == "" {
		return testArgs{}, fmt.Errorf("отсутствует AMOUNT")
	}

	password := os.Getenv("METAMASK_PASSWORD")
	if password == "" {
		return testArgs{}, fmt.Errorf("отсутствует METAMASK_PASSWORD")
	}

	return testArgs{
		seed:     seed,
		token1:   token1,
		token2:   token2,
		amount:   amount,
		password: password,
	}, nil
}
