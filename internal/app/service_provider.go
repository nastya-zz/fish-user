package app

import (
	"context"
	"log"
	"user/internal/api/user"
	"user/internal/client/broker"
	"user/internal/client/broker/rabbitmq"
	"user/internal/client/db"
	"user/internal/client/db/pg"
	"user/internal/client/minio/minio"
	"user/internal/closer"
	"user/internal/config"
	"user/internal/consumer"
	rmqConsumer "user/internal/consumer/rabbitmq"
	"user/internal/repository"
	"user/internal/repository/settings"
	"user/internal/repository/subscriptions"
	userRepository "user/internal/repository/user"
	"user/internal/service"
	"user/internal/service/event"
	settingsService "user/internal/service/settings"
	sbService "user/internal/service/subscribtions"
	userService "user/internal/service/user"
	"user/internal/transaction"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	rmqConfig   config.RMQConfig
	minioConfig *config.MinioConfig

	rmqClient   broker.ClientMsgBroker
	dbClient    db.Client
	txManager   db.TxManager
	minioClient *minio.Client

	eventConsumer consumer.Consumer

	userRepository          repository.UserRepository
	settingsRepository      repository.SettingsRepository
	subscriptionsRepository repository.SubscriptionsRepository

	userService          service.UserService
	settingsService      service.SettingsService
	subscriptionsService service.SubscriptionsService
	eventService         service.EventsService

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

func (s *serviceProvider) RMQConfig() config.RMQConfig {
	if s.rmqConfig == nil {
		cfg, err := config.NewRMQConfig()
		if err != nil {
			log.Fatalf("failed to get rmqConfig : %s", err.Error())
		}

		s.rmqConfig = cfg
	}

	return s.rmqConfig
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

func (s *serviceProvider) MinioConfig() *config.MinioConfig {
	if s.minioConfig == nil {
		cfg, err := config.NewMinioConfig()
		if err != nil {
			log.Fatalf("failed to get minio config: %s", err.Error())
		}

		s.minioConfig = cfg
	}

	return s.minioConfig
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

func (s *serviceProvider) RabbitMQClient(ctx context.Context) broker.ClientMsgBroker {
	if s.rmqClient == nil {
		cl, err := rabbitmq.NewRabbitMQ(ctx, s.RMQConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create rmq client: %v", err)
		}

		closer.Add(cl.Close)

		s.rmqClient = cl
	}

	return s.rmqClient
}

func (s *serviceProvider) MinioClient(ctx context.Context) *minio.Client {
	if s.rmqClient == nil {
		cl, err := minio.New(ctx, s.minioConfig.Endpoint, s.minioConfig.AccessKey, s.minioConfig.SecretKey)
		if err != nil {
			log.Fatalf("failed to create minio client: %v", err)
		}

		s.minioClient = cl
	}

	return s.minioClient
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
			s.SettingsService(ctx),
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

func (s *serviceProvider) EventService(ctx context.Context) service.EventsService {
	if s.eventService == nil {
		s.eventService = event.New(
			s.UserService(ctx),
		)
	}

	return s.eventService
}

func (s *serviceProvider) EventConsumer(ctx context.Context) consumer.Consumer {
	if s.eventConsumer == nil {
		r := s.RabbitMQClient(ctx)
		s.eventConsumer = rmqConsumer.NewUserConsumer(r.Connect().Channel, s.EventService(ctx))
	}
	return s.eventConsumer
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {

	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(
			s.UserService(ctx),
			s.SettingsService(ctx),
			s.SubscriptionsService(ctx),
		)
	}

	return s.userImpl
}
