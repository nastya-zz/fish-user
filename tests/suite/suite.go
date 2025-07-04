package suite

import (
	"context"
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	desc "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"user/internal/client/db"
	"user/internal/config"
	"user/pkg/logger"
)

type Suite struct {
	*testing.T                   // Потребуется для вызова методов *testing.T внутри Suite
	Cfg        config.GRPCConfig // Конфигурация grpc
	UserClient desc.UserV1Client // Клиент для взаимодействия с gRPC-сервером
	dbClient   db.Client
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()

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
		logger.Fatal("failed to create db client", "error", err)
	}

	tx, err := cl.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback() // Откат после теста

	if _, err = cl.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		logger.Fatal("failed to create uuid extension", "error", err)
	}
	fixtures, err := testfixtures.New(
		testfixtures.Database(cl),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("/Users/nastya/Documents/GitHub/fish-services/user/tests/fixtures"), // The directory containing the YAML files
	)

	if err != nil {
		logger.Fatal("failed to create fixtures", "error", err, "dsn", dbConfig.DSN())
	}

	if err = fixtures.Load(); err != nil {
		logger.Fatal("failed to load fixtures", "error", err)
	}

	t.Cleanup(func() {
		t.Helper()
		err = cleanTables(cl)
		if err != nil {
			t.Fatalf("failed to drop data in tables: %v", err)
		}
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

func cleanTables(cl *sql.DB) error {
	cleanupQuery := `
					TRUNCATE TABLE users CASCADE;
					TRUNCATE TABLE settings CASCADE;
					TRUNCATE TABLE follows CASCADE;
					TRUNCATE TABLE user_blocks CASCADE;
					`
	if _, err := cl.Exec(cleanupQuery); err != nil {
		return err
	}
	return nil
}
