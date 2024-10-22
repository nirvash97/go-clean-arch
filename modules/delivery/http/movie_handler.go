package http

import (
	"encoding/json"
	"go-clean-arch/modules/usecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	uc *usecase.MovieUsecase
}

func NewMovieHandler(uc *usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{uc: uc}
}

func (h *MovieHandler) GetMovieBylanguage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lang := params["language"]

	movies, err := h.uc.GetMovieByLanguage(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(movies) <= 0 {
		http.Error(w, "no item found", http.StatusNoContent)
		return
	}

	decodeData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(movies)
	w.Header().Set("Content-Type", "application/json")
	w.Write(decodeData)
}
