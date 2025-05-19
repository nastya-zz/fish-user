package main

import (
	"context"
	"log"
	"user/internal/app"
)

/***
TODO:
- использовать логгер
- интерсептор на авторизацию

- хелс чек сервиса

- ТЕСТЫ функциональные
*/

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
