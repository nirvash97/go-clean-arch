package http

import (
	"encoding/json"
	"fmt"
	"go-clean-arch/modules/entities/movies"
	"go-clean-arch/modules/usecase"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

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
func (h *MovieHandler) GetMovieByLanguagePagination(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lang := params["language"]
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("perPage")
	log.Println("Param page ::: " + page)
	log.Println("param PerPage ::: " + perPage)

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	perPageInt, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var regexLetter = regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regexLetter.MatchString(lang) {
		http.Error(w, "Parameter must contain only word", http.StatusBadRequest)
		return
	}
	// --------------------- Count Data ------------

	itemCount, err := h.uc.GetMovieByLanguageItemCount(lang)
	if err != nil {
		const timeFormat = time.RFC3339
		fmt.Printf("[%s] %s %s \n", time.Now().UTC().Format(timeFormat), r.Method, r.URL)
	}
	log.Println("item count ::: " + fmt.Sprintf("%d", itemCount))

	// --------- Main Data-------------
	movieList, err := h.uc.GetMovieByLanguagePagination(lang, pageInt, perPageInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(movieList) <= 0 {
		http.Error(w, "", http.StatusNoContent)
		return
	}

	pattern := movies.MoviePagination{Page: pageInt, PerPage: perPageInt, TotalRow: itemCount, Data: movieList}
	decodeData, err := json.Marshal(pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(decodeData)

}

// func (h *MovieHandler) getMovieTotalRecord(w http.ResponseWriter, r *http.Response) {

// }
