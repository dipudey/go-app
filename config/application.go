package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/dipudey/go-app/internal/router"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	GinRouter *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	GinRouter = gin.New()
}

func RunApplication() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx, _ := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	// Connect Database
	dbConnection, err := cfg.Database.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db, err := dbConnection.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}
	defer db.Close()

	GinRouter = router.InitRoutes(dbConnection)

	log.Printf("Starting %s server...", cfg.AppName)
	log.Printf("Server running on http://%s:%d", cfg.Server.Host, cfg.Server.Port)
	s := &http.Server{
		Addr:           ":" + fmt.Sprintf("%d", cfg.Server.Port),
		Handler:        GinRouter,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()
	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("server stopped gracefully")
}
