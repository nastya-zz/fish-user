package suite

import (
	"context"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"user/internal/config"
)

type Suite struct {
	*testing.T                   // Потребуется для вызова методов *testing.T внутри Suite
	Cfg        config.GRPCConfig // Конфигурация grpc
	UserClient desc.UserV1Client // Клиент для взаимодействия с gRPC-сервером
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	err := config.Load("../../.env")
	if err != nil {
		t.Fatalf("grpc server read config failed: %v", err)
	}
	cfg, err := config.NewGRPCConfig()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.Timeout())

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(context.Background(),
		cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials())) // Используем insecure-коннект для тестов
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		UserClient: desc.NewUserV1Client(cc),
	}
}
