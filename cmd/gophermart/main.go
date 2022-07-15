package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/businesslogic"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/httpserver/handlers"
)

func main() {
	//
	// init service
	// load config service
	// run main gorutines
	//

	waitstop := 1

	config.ConfigCLS.LoadConfig()
	defer config.LoggerCLS.Info("CLS server stopped.")
	defer config.LoggerCLS.Sync()
	//
	// final defer insert here
	//
	defer config.LoggerCLS.Info("CLS server begin stop...")
	config.LoggerCLS.Info("CLS server start")

	//
	// run main gorutines
	//

	//
	// run http server
	//
	handlers.GoListenRutine()
	config.LoggerCLS.Info("CLS server http listener started on " + config.ConfigCLS.RunAddress)

	// run update orders statuses queue
	bs := new(businesslogic.BusinessSession)
	bs.UpdateAllOrdersFromAccrual(time.Second * 1)

	config.LoggerCLS.Info("CLS server update order statuses started")
	// wait signals
	// we need to reserve to buffer size 1, so the notifier are not blocked
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	config.LoggerCLS.Info("CLS server begin waiting signal...")
	sig := <-chanOS

	config.LoggerCLS.Sugar().Infof("INFO got a signal '%v', start shutting down... wait %v seconds\n", sig, waitstop)
	<-time.After(time.Second * time.Duration(waitstop))

}
