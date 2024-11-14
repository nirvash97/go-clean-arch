package service

import (
	"go-clean-arch/modules/database/client"
	"go-clean-arch/modules/delivery/http"
	"go-clean-arch/modules/repositories/authrepo"
	movierepo "go-clean-arch/modules/repositories/movie"
	"go-clean-arch/modules/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

func InitService() *echo.Echo {
	e := echo.New()

	client, err := client.InitialServer()

	// defer func() {
	// 	if err = client.Disconnect(context.Background()); err != nil {
	// 		log.Fatalf("Connect to database failed : %v", err)
	// 	}
	// }()
	if err != nil {
		log.Fatalf("Connect to database failed : %v", err)
	}

	movieRepo := movierepo.NewMovieMongoRepo(client)
	movieUseCase := usecase.NewMovieUsecase(movieRepo)
	http.NewMovieHandler(e, movieUseCase)

	authRepo := authrepo.NewAuthMongoRepo(client)
	authuseCase := usecase.NewAuthUsecase(authRepo)
	http.NewAuthHandler(e, authuseCase)

	return e
}
