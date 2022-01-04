package main

import (
	"context"
	"flag"
	"fmt"
	config "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
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

	conf, err := config.NewConfig(configFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	logg := logger.New(conf.Logger)

	switch conf.App.Storage {
	case "memory":
		storage = memorystorage.New()
	case "sql":
		storage = sqlstorage.New(conf)
	default:
		fmt.Println("no storage found. Memory storage run by default")
		storage = memorystorage.New()
	}

	calendar := app.New(logg, storage, &conf.App)

	server := internalhttp.NewServer(logg, calendar, &conf.App)

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
