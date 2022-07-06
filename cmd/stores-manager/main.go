package main

import (
	"context"
	"fmt"
	"github.com/igorralexsander/stores-manager/internal"
	"github.com/igorralexsander/stores-manager/internal/data/module"
	serviceModule "github.com/igorralexsander/stores-manager/internal/domain/module"
	"github.com/igorralexsander/stores-manager/internal/infra/config"
	infraModule "github.com/igorralexsander/stores-manager/internal/infra/module"
	"github.com/igorralexsander/stores-manager/internal/infra/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	clientsMod := module.NewClientsModule()
	scyllaClient := clientsMod.ProvideScyllaClient(config.Instance().GetDatabaseScyllaConfig())
	if err := scyllaClient.Connect(); err != nil {
		log.Error(err, "Failed to connect in scylladb")
	}

	repositoryMod := module.NewRepoistoryModule(scyllaClient)
	servicesMod := serviceModule.NewServiceModule()
	restMod := infraModule.NewRestModule()

	application := internal.NewApplication(servicesMod, restMod, clientsMod, repositoryMod)

	apiServer := rest.NewServer(application)

	httpServer := apiServer.CreateHttpServer()

	go apiServer.Start(httpServer, config.Instance().GetServerConfig().Host)

	shutDownHook(httpServer)

}

func shutDownHook(apiServer *echo.Echo) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	switch <-quit {
	case os.Interrupt, syscall.SIGTERM:

		log.Info("Initialize Gracefully shutdown")

		log.Info(fmt.Sprintf("Wait %d seconds to process pending requests", 10))
		time.Sleep(10 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		log.Info("Shutdown HTTP server...")
		if err := apiServer.Shutdown(ctx); err != nil {
			log.Fatal(err, "Error to gracefully stop application, application stopped.")
		}
		log.Info("Complete Gracefully shutdown")
	}

}
