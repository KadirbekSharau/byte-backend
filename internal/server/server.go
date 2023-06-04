package server

import (
	"log"
	"net/http"

	"github.com/KadirbekSharau/Byte/configs"
	"github.com/KadirbekSharau/Byte/internal/auth"
	authRepository "github.com/KadirbekSharau/Byte/internal/auth/repository"
	authUsecase "github.com/KadirbekSharau/Byte/internal/auth/usecase"
)

type Server struct {
	httpServer *http.Server
	authUC auth.UseCase
}

func NewServer() *Server {
	db, err := configs.NewPostgresDB()
	if err != nil {
		log.Fatal("Can't connect to the database. Error: %s", err.Error())
	}

	authRepository := authRepository.NewUserRepository(db)
	return &Server{
		authUC: authUsecase.NewAuthUseCase(authRepository, "", []byte{}, 12),
	}
}

func (*Server) Run() {

}
