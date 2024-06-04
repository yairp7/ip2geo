package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yairp7/ip2geo/internal/common"
)

var loggerImpl common.Logger = common.NewStdoutLogger(common.DEBUG)

func Start() {
	err := godotenv.Load()
	if err != nil {
		loggerImpl.Warn(".env file not found\n")
	}

	eEnv := os.Getenv("ENV")
	ePort := os.Getenv("PORT")

	if len(eEnv) == 0 || len(ePort) == 0 {
		panic(fmt.Errorf("invalid enviroment variables - ENV=%s, PORT=%s", eEnv, ePort))
	}

	port, err := strconv.Atoi(ePort)
	if err != nil {
		panic(fmt.Errorf("bad port - %v", err))
	}

	if eEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := NewRouter(loggerImpl)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "", port),
		Handler: router,
	}

	ctx, cancelSignalNotify := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelSignalNotify()

	loggerImpl.Info("Server ready on port %d\n", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("server error: %s\n", err))
		}
	}()

	<-ctx.Done()
	shutdown(srv)
}

func shutdown(srv *http.Server) {
	loggerImpl.Info("Shutting server down...")

	ShutdownRouter()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Errorf("Server shutdown failed: %s\n", err))
	}

	select {
	case <-ctx.Done():
		loggerImpl.Warn("Server shutdown timeout")
	default:
		loggerImpl.Info("Server shutdown complete")
	}
}
