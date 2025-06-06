package app

import (
	"context"
	"flag"
	descAuth "github.com/nastya-zz/fisher-protocols/gen/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
	"user/internal/closer"
	"user/internal/config"
	"user/internal/logger"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

const (
	envTest = "test"
	envDev  = "dev"
	envProd = "prod"
)

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	a.runEventConsumer(ctx)

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	a.serviceProvider.LoggerConfig()
	a.setupLogger()

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	path := a.mustPath()
	err := config.Load(path)

	log.Println("initConfig with ", "path: ", path)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	if env := a.serviceProvider.loggerConfig.Environment(); env == envTest {
		a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	} else {
		a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	}

	reflection.Register(a.grpcServer)

	descAuth.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Info("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runEventConsumer(ctx context.Context) {
	a.serviceProvider.EventConsumer(ctx)
	a.serviceProvider.eventConsumer.Start(ctx)
}

func (a *App) setupLogger() {
	env := a.serviceProvider.loggerConfig.Environment()

	switch env {
	case envTest:
		logger.Init(slog.LevelDebug)

	case envDev:
		logger.Init(slog.LevelDebug)

	case envProd:
		logger.Init(slog.LevelInfo)

	default: // If env config is invalid, set prod settings by default due to security
		logger.Init(slog.LevelInfo)
	}
}

func (a *App) mustPath() string {
	env := flag.String(
		"env",
		"dev",
		"environment",
	)
	flag.Parse()

	if *env == "" {
		*env = envDev
	}

	switch *env {
	case envTest:
		return ".env.test"
	case envDev:
		return ".env"
	case envProd:
		return ".env.prod"
	default:
		return ".env"
	}
}
