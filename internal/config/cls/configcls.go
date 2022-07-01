package config

type Config struct {
	RunAddress           string
	DataBaseURI          string
	AccrualSystemAddress string
}

func (c *Config) LoadConfig() (err error) {

	//
	// TODO: load config from environment and comman flags
	// адрес и порт запуска сервиса: переменная окружения ОС RUN_ADDRESS или флаг -a;
	// адрес подключения к базе данных: переменная окружения ОС DATABASE_URI или флаг -d;
	// адрес системы расчёта начислений: переменная окружения ОС ACCRUAL_SYSTEM_ADDRESS или флаг -r.

	return nil
}
