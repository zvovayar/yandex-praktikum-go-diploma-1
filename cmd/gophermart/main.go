package main

import config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"

func main() {
	//
	// TODO:
	// init service
	// load config service
	// run main gorutines
	//

	config.ConfigCLS.LoadConfig()
	defer config.LoggerCLS.Sync()
	config.LoggerCLS.Info("CLS server start")

	//
	// TODO: run main gorutines
	//
	config.LoggerCLS.Info("CLS server started")

	//
	// TODO: run http server
	//

}
