package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wenkaler/blackbox/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/wenkaler/blackbox/storage"
)

const serviceName = "blackbox"

var serviceVersion = "dev"

type configuration struct {
	Addr  string `envconfig:"ADDR" default:":8080"`
	Debug bool   `envconfig:"DEBUG" default:"false"`

	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	DBName   string `envconfig:"DB_DBNAME" default:"blackbox"`
	SSLmode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(serviceVersion)
		os.Exit(0)
	}
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	var cfg configuration
	if err := envconfig.Process("BLACKBOX", &cfg); err != nil {
		level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}

	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	storage, err := storage.New(&storage.Config{
		Logger:      logger,
		ServiceName: serviceName,

		User:     cfg.User,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		DBName:   cfg.DBName,
		SSLmode:  cfg.SSLmode,
	})
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize storage", "err", err)
		os.Exit(1)
	}

	server, err := server.New(&server.Config{
		Logger:  logger,
		Addr:    cfg.Addr,
		Storage: storage,
	})
	if err != nil {
		level.Error(logger).Log("msg", "failed to initialize server", "err", err)
		os.Exit(1)
	}
	go func() {
		level.Info(logger).Log("msg", "starting http server", "addr", cfg.Addr)
		if err := server.Run(); err != nil {
			level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c
	level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

	if err := server.Shutdown(); err != nil {
		level.Error(logger).Log("msg", "server shutdown failure", "err", err)
	}

	storage.Shutdown()
	level.Info(logger).Log("msg", "goodbye")

}
