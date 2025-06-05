package suite

import (
	"context"
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"user/internal/client/db"
	"user/internal/config"
)

type Suite struct {
	*testing.T                   // Потребуется для вызова методов *testing.T внутри Suite
	Cfg        config.GRPCConfig // Конфигурация grpc
	UserClient desc.UserV1Client // Клиент для взаимодействия с gRPC-сервером
	//fixtures   *testfixtures.Loader
	dbClient db.Client
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	//t.Parallel()

	err := config.Load("../../.env.test")

	if err != nil {
		t.Fatalf("grpc server read config failed: %v", err)
	}
	cfg, err := config.NewGRPCConfig()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.Timeout())

	dbConfig, err := config.NewPGConfig()
	if err != nil {
		t.Fatal(err)
	}

	cl, err := sql.Open("postgres", dbConfig.DSN())

	if err != nil {
		log.Fatalf("failed to create db client: %v", err)
	}

	tx, err := cl.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback() // Откат после теста

	if _, err = cl.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		log.Fatalf("failed to create uuid extension: %v", err)
	}
	fixtures, err := testfixtures.New(
		testfixtures.Database(cl),        // You database connection
		testfixtures.Dialect("postgres"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("/Users/nastya/Documents/GitHub/fish-services/user/tests/fixtures"), // The directory containing the YAML files
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	if err != nil {
		log.Fatalf("failed to create fixtures: %v, %s", err, dbConfig.DSN())
	}

	if err = fixtures.Load(); err != nil {
		log.Fatalf("failed to load fixtures: %v", err)
	}

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
