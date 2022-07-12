package config

import (
	"flag"

	"github.com/caarlos0/env"

	"go.uber.org/zap"
)

// Global variables
var LoggerCLS *zap.Logger
var ConfigCLS Config

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	DataBaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DebugLogger          string
}

func (c *Config) LoadConfig() (err error) {

	//
	// load config from environment and comman flags
	// адрес и порт запуска сервиса: переменная окружения ОС RUN_ADDRESS или флаг -a;
	// адрес подключения к базе данных: переменная окружения ОС DATABASE_URI или флаг -d;
	// адрес системы расчёта начислений: переменная окружения ОС ACCRUAL_SYSTEM_ADDRESS или флаг -r.

	// Default config init
	ConfigCLS.RunAddress = "localhost:8888"
	ConfigCLS.DataBaseURI = ""
	ConfigCLS.AccrualSystemAddress = "localhost:8080"
	ConfigCLS.DebugLogger = "+"

	// load flags
	cflags := new(Config)
	flag.StringVar(&cflags.RunAddress, "a", "", "address to listen on")
	flag.StringVar(&cflags.DataBaseURI, "d", "", "database URI")
	flag.StringVar(&cflags.AccrualSystemAddress, "r", "", "accrual system address")
	flag.StringVar(&cflags.DebugLogger, "v", "+", "switch off debug logger (-)")
	flag.Parse()

	if cflags.DebugLogger == "+" {
		LoggerCLS.Sync()
		LoggerCLS, err = zap.NewDevelopment()
		if err != nil {
			LoggerCLS.Panic("can't create zap developmetn logger")
		}
	}

	// load environment variables
	err = env.Parse(c)
	if err != nil {
		LoggerCLS.Fatal(err.Error())
		return err
	}

	LoggerCLS.Debug("loaded environment variables:")
	LoggerCLS.Sugar().Debug(*c)

	LoggerCLS.Debug("loaded flags:")
	LoggerCLS.Sugar().Debug(*cflags)

	if cflags.RunAddress != "" {
		ConfigCLS.RunAddress = cflags.RunAddress
	} else if c.RunAddress != "" {
		ConfigCLS.RunAddress = c.RunAddress
	}
	if cflags.DataBaseURI != "" {
		ConfigCLS.DataBaseURI = cflags.DataBaseURI
	} else if c.DataBaseURI != "" {
		ConfigCLS.DataBaseURI = c.DataBaseURI
	}
	if cflags.AccrualSystemAddress != "" {
		ConfigCLS.AccrualSystemAddress = cflags.AccrualSystemAddress
	} else if c.AccrualSystemAddress != "" {
		ConfigCLS.AccrualSystemAddress = c.AccrualSystemAddress
	}
	if cflags.DebugLogger != "" {
		ConfigCLS.DebugLogger = cflags.DebugLogger
	} else if c.DebugLogger != "" {
		ConfigCLS.DebugLogger = c.DebugLogger
	}

	LoggerCLS.Info("effective config variables:")
	LoggerCLS.Sugar().Info(ConfigCLS)

	return err
}

func init() {

	var err error

	LoggerCLS, err = zap.NewProduction()

	if err != nil {
		LoggerCLS.Panic("can't create zap production logger")
	}
}
