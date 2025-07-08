package main

import (
	"context"
	"user/internal/app"
	"user/pkg/logger"
)

/***
TODO:
- интерсептор на авторизацию

- хелс чек сервиса
- минио и рабит вынести отдельно


- ТЕСТЫ функциональные
*/

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatal("failed to init app", "error", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		logger.Fatal("failed to run app", "error", err.Error())
	}
}
