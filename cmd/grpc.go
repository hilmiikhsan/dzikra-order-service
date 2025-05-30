package cmd

import (
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/cmd/proto/order"
	"github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/cmd/proto/transaction"
	"github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/internal/adapter"
	"github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/internal/infrastructure"
	"github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/internal/infrastructure/config"
	orderService "github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/internal/module/order/handler/grpc"
	transactionService "github.com/Digitalkeun-Creative/be-dzikra-ecommerce-order-service/internal/module/transaction/handler/grpc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func RunServeGRPC() {
	envs := config.Envs

	logLevel, err := zerolog.ParseLevel(envs.App.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	infrastructure.InitializeLogger(
		envs.App.Environtment,
		envs.App.LogFile,
		logLevel,
	)

	lis, err := net.Listen("tcp", ":"+envs.App.GrpcPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen on gRPC port")
	}

	grpcServer := grpc.NewServer()

	opts := []adapter.Option{
		adapter.WithGRPCServer(grpcServer),
		adapter.WithDzikraPostgres(),
		adapter.WithDzikraRedis(),
		adapter.WithDzikraMidtrans(),
		adapter.WithRabbitMQ(),
	}

	// Sync adapters
	if err := adapter.Adapters.Sync(opts...); err != nil {
		log.Fatal().Err(err).Msg("Failed to sync adapters")
	}

	order.RegisterOrderServiceServer(grpcServer, orderService.NewOrderAPI())
	transaction.RegisterTransactionServiceServer(grpcServer, transactionService.NewTransactionAPI())

	go func() {
		log.Info().Msgf("gRPC server is running on port %s", envs.App.GrpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("Failed to serve gRPC")
		}
	}()

	quit := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	if runtime.GOOS == "windows" {
		signals = []os.Signal{os.Interrupt}
	}
	signal.Notify(quit, signals...)
	<-quit

	log.Info().Msg("gRPC server is shutting down ...")
	grpcServer.GracefulStop()

	if err := adapter.Adapters.Unsync(); err != nil {
		log.Error().Err(err).Msg("Error while closing adapters")
	}

	log.Info().Msg("gRPC server gracefully stopped")
}
