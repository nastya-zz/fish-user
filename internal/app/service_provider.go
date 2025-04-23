package app

import (
	"context"
	"log"
	"user/internal/api/user"
	"user/internal/client/db"
	"user/internal/client/db/pg"
	"user/internal/closer"
	"user/internal/config"
	"user/internal/repository"
	"user/internal/repository/settings"
	"user/internal/repository/subscriptions"
	userRepository "user/internal/repository/user"
	"user/internal/service"
	settingsService "user/internal/service/settings"
	sbService "user/internal/service/subscribtions"
	userService "user/internal/service/user"
	"user/internal/transaction"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository          repository.UserRepository
	settingsRepository      repository.SettingsRepository
	subscriptionsRepository repository.SubscriptionsRepository

	userService          service.UserService
	settingsService      service.SettingsService
	subscriptionsService service.SubscriptionsService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}
func (s *serviceProvider) SettingsRepository(ctx context.Context) repository.SettingsRepository {
	if s.settingsRepository == nil {
		s.settingsRepository = settings.NewRepository(s.DBClient(ctx))
	}

	return s.settingsRepository
}
func (s *serviceProvider) SubscriptionsRepository(ctx context.Context) repository.SubscriptionsRepository {
	if s.subscriptionsRepository == nil {
		s.subscriptionsRepository = subscriptions.NewRepository(s.DBClient(ctx))
	}

	return s.subscriptionsRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}
func (s *serviceProvider) SettingsService(ctx context.Context) service.SettingsService {
	if s.settingsService == nil {
		s.settingsService = settingsService.NewService(
			s.SettingsRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.settingsService
}
func (s *serviceProvider) SubscriptionsService(ctx context.Context) service.SubscriptionsService {
	if s.subscriptionsService == nil {
		s.subscriptionsService = sbService.NewService(
			s.SubscriptionsRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.subscriptionsService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *user.Implementation {

	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(
			s.UserService(ctx),
			s.SettingsRepository(ctx),
			s.SubscriptionsRepository(ctx),
		)
	}

	return s.userImpl
}
