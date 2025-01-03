package routes

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
*
* factory
* <p>
* factory file
*
* Copyright (c) 2025 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 2/01/2025
*
 */

type Port int64
type Name string

const InstanceRouterGin int = iota
const InstanceRouterGorilla = 1
const NameRouterGin string = "Gin"
const NameRouterGorilla string = "Gorilla Mux"

type ServerFactory struct {
	*zap.Logger
}

type IServerFactory interface {
	Run()
}

func NewServerFactory(logger *zap.Logger, instance int, port Port, name Name) (IServerFactory, error) {
	switch instance {
	case InstanceRouterGin:
		return newGinFramework(logger, port, name), nil
	default:
		return nil, errors.New("invalid instance")

	}
}

// SetupSignalHandler configura el manejo de señales para una parada controlada.
func setupSignalHandler(log *zap.Logger) (quitOs <-chan struct{}) {

	quit := make(chan struct{})
	// Canal para recibir señales del sistema
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Espera la primera señal y cierra el canal `stop`.
		next := <-s
		log.Info("Caught signal; shutting down...", zap.Any("signal", next))
		close(quit)
		// Espera una segunda señal para terminar inmediatamente.
		next = <-s
		log.Info("Caught signal next; shutting down...", zap.Any("signal", next))
		os.Exit(1)
	}()
	return quit
}

func createHttpServer(port Port, middleware *negroni.Negroni) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware,
	}
	return srv
}

func listenAndServe(port Port, name Name, middleware *negroni.Negroni, log *zap.Logger) {
	srv := createHttpServer(port, middleware)
	stop := setupSignalHandler(log)

	go func() {
		log.Info(fmt.Sprintf("Starting http server on port %d [%s]...", port, name))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error starting http server", zap.Error(err))
			return
		}
	}()

	// Esperar a recibir una señal
	<-stop
	log.Info("shutting down http server", zap.Int("port", int(port)))
	// Establece un tiempo límite para la parada del servidor.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	// Intenta cerrar el servidor de manera ordenada.
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("error shutting down http server", zap.Error(err))
	}

	log.Info("shutting down http server", zap.Int("port", int(port)))
}
