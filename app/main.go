package main

import (
	"context"
	"go-clean-arch/modules/database/client"
	handler "go-clean-arch/modules/delivery/http"
	"go-clean-arch/modules/repositories"
	"go-clean-arch/modules/usecase"
	middleware "go-clean-arch/pkg/middlewares"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	client, err := client.InitialServer()
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("sample_mflix")
	// Initialize repository, use case, and handler
	movieRepo := repositories.NewMovieMongoRepo(db)
	movieUseCase := usecase.NewMovieUsecase(movieRepo)
	movieHandler := handler.NewMovieHandler(movieUseCase)

	// Auth Service
	authRepo := repositories.NewAuthMongoRepo(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Setup router

	r := mux.NewRouter()

	// Define Route
	go r.HandleFunc("/movie/languages/{language}", movieHandler.GetMovieBylanguage).Methods("GET")
	go r.HandleFunc("/movie/language/pagination/{language}", movieHandler.GetMovieByLanguagePagination).Methods("GET")
	go r.HandleFunc("/signUp", authHandler.HandleSignUp).Methods("POST")
	go r.HandleFunc("/signIn", authHandler.HandleAuth).Methods("POST")
	log.Println("Server is listening on port :: 8081 ")
	server := &http.Server{
		Addr:         ":8081",
		Handler:      middleware.Middleware(r),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	errServer := server.ListenAndServe()
	if errServer != nil {
		log.Println(errServer)
	}
}
