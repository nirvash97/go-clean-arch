package http

import (
	"encoding/json"
	"go-clean-arch/modules/usecase"
	"net/http"
	"regexp"

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
	// Check param should have only letter [return 400]
	var regex = regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(lang) {
		http.Error(w, "Parameter must contain only word", http.StatusBadRequest)
		return
	}

	movies, err := h.uc.GetMovieByLanguage(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Check is empty [return 204]
	if len(movies) <= 0 {
		http.Error(w, "no item found", http.StatusNoContent)
		return
	}

	decodeData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(decodeData)
}
