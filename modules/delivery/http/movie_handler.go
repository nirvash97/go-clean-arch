package http

import (
	"fmt"
	"go-clean-arch/modules/entities/movies"
	"go-clean-arch/modules/usecase"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	uc *usecase.MovieUsecase
}

func NewMovieHandler(e *echo.Echo, uc *usecase.MovieUsecase) *MovieHandler {
	handler := &MovieHandler{uc: uc}
	e.GET("/movie/test", handler.GetMoviesTest)
	e.GET("/movie/language/:language", handler.GetEchoMovieBylanguage)
	e.GET("/movie/language/pagination/:language", handler.GetEchoMovieByLanguagePagination)
	return handler
}

// ================= ECHO ===================
func (h *MovieHandler) GetMoviesTest(c echo.Context) error {
	movies, err := h.uc.GetMoviesTest()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetEchoMovieBylanguage(c echo.Context) error {
	lang := c.Param("language")
	// Check param should have only letter [return 400]
	var regex = regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(lang) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Parameter language must contain only word",
		})

	}
	movies, err := h.uc.GetMovieByLanguage(lang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Check is empty [return 204]
	if len(movies) <= 0 {

		return c.JSON(http.StatusNoContent, map[string]string{
			"message": "no item found",
		})
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, movies)

}

func (h *MovieHandler) GetEchoMovieByLanguagePagination(c echo.Context) error {
	lang := c.Param("language")
	page := c.Request().URL.Query().Get("page")
	perPage := c.Request().URL.Query().Get("perPage")
	log.Println("Param page ::: " + page)
	log.Println("param PerPage ::: " + perPage)
	var regex = regexp.MustCompile(`^[a-zA-Z]+$`)
	pageInt, err := strconv.ParseInt(page, 10, 64)

	if !regex.MatchString(lang) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Parameter language must contain only word",
		})

	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Query Param 'page' is required",
		})

	}
	perPageInt, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Query Param 'perPage' is required",
		})
	}
	var regexLetter = regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regexLetter.MatchString(lang) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Parameter must contain only word",
		})
	}
	// --------------------- Count Data ------------

	itemCount, err := h.uc.GetMovieByLanguageItemCount(lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "get movie pagination item count failed",
		})
	}

	log.Println("item count ::: " + fmt.Sprintf("%d", itemCount))

	// --------- Main Data-------------
	movieList, err := h.uc.GetMovieByLanguagePagination(lang, pageInt, perPageInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if len(movieList) <= 0 {
		return c.JSON(http.StatusNoContent, map[string]string{
			"message": "",
		})
	}
	pattern := movies.MoviePagination{
		Page:     pageInt,
		PerPage:  perPageInt,
		TotalRow: itemCount,
		Data:     movieList,
	}
	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, pattern)
}

// // ================== MUX Gorrila ====================

// func (h *MovieHandler) GetMovieBylanguage(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	lang := params["language"]
// 	// Check param should have only letter [return 400]
// 	var regex = regexp.MustCompile(`^[a-zA-Z]+$`)
// 	if !regex.MatchString(lang) {
// 		http.Error(w, "Parameter must contain only word", http.StatusBadRequest)
// 		return
// 	}

// 	movies, err := h.uc.GetMovieByLanguage(lang)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	// Check is empty [return 204]
// 	if len(movies) <= 0 {
// 		http.Error(w, "no item found", http.StatusNoContent)
// 		return
// 	}

// 	decodeData, err := json.Marshal(movies)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(decodeData)
// }
// func (h *MovieHandler) GetMovieByLanguagePagination(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	lang := params["language"]
// 	page := r.URL.Query().Get("page")
// 	perPage := r.URL.Query().Get("perPage")
// 	log.Println("Param page ::: " + page)
// 	log.Println("param PerPage ::: " + perPage)

// 	pageInt, err := strconv.ParseInt(page, 10, 64)

// 	if err != nil {
// 		http.Error(w, "Query Param 'page' is required", http.StatusBadRequest)
// 		return
// 	}

// 	perPageInt, err := strconv.ParseInt(perPage, 10, 64)
// 	if err != nil {
// 		http.Error(w, "Query Param 'perPage' is required", http.StatusBadRequest)
// 		return
// 	}
// 	var regexLetter = regexp.MustCompile(`^[a-zA-Z]+$`)
// 	if !regexLetter.MatchString(lang) {
// 		http.Error(w, "Parameter must contain only word", http.StatusBadRequest)
// 		return
// 	}
// 	// --------------------- Count Data ------------

// 	itemCount, err := h.uc.GetMovieByLanguageItemCount(lang)
// 	if err != nil {
// 		const timeFormat = time.RFC3339
// 		fmt.Printf("[%s] %s %s \n", time.Now().UTC().Format(timeFormat), r.Method, r.URL)
// 	}
// 	log.Println("item count ::: " + fmt.Sprintf("%d", itemCount))

// 	// --------- Main Data-------------
// 	movieList, err := h.uc.GetMovieByLanguagePagination(lang, pageInt, perPageInt)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if len(movieList) <= 0 {
// 		http.Error(w, "", http.StatusNoContent)
// 		return
// 	}

// 	pattern := movies.MoviePagination{Page: pageInt, PerPage: perPageInt, TotalRow: itemCount, Data: movieList}
// 	decodeData, err := json.Marshal(pattern)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return

// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(decodeData)

// }

// // func (h *MovieHandler) getMovieTotalRecord(w http.ResponseWriter, r *http.Response) {

// // }
