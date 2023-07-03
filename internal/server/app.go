package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KadirbekSharau/Byte/config"
	"github.com/KadirbekSharau/Byte/internal/auth"
	authhttp "github.com/KadirbekSharau/Byte/internal/auth/delivery/http"
	authRepository "github.com/KadirbekSharau/Byte/internal/auth/repository"
	authUsecase "github.com/KadirbekSharau/Byte/internal/auth/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
}

func NewApp() *App {
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal("Can't connect to the database. Error: %s", err.Error())
	}

	authRepository := authRepository.NewUserRepository(db)
	return &App{
		authUC: authUsecase.NewAuthUseCase(
			authRepository,
			os.Getenv("HASH_SALT"),
			[]byte(os.Getenv("SIGNING_KEY")),
			viper.GetDuration("auth.token_ttl"),
		),
	}
}

func (a *App) Run(port string) error {
	// Init Gin Handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
