package main

import (
	"context"
	"flag"
	config2 "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	sqlstorage "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func main() {
	var storage app.Storage
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := config2.NewConfig(configFile)
	logg := logger.New(config.Logger.Level)

	switch config.App.Storage {
	case "memory":
		storage = memorystorage.New()
	case "sql":
		storage = sqlstorage.New(config)
	default:
		storage = memorystorage.New()
	}

	calendar := app.New(logg, storage, &config.App)

	server := internalhttp.NewServer(logg, calendar, &config.App)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
