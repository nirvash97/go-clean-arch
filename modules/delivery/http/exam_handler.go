package http

import (
	"go-clean-arch/modules/entities/exam"
	"go-clean-arch/modules/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ExamHandler struct {
	uc *usecase.ExamUseCase
}

func NewExamHandler(e *echo.Echo, uc *usecase.ExamUseCase) *ExamHandler {
	handler := &ExamHandler{uc: uc}
	e.GET("/exam/hello", handler.getHelloWorld)
	e.GET("/exam/user-list", handler.getAlluser)
	e.POST("/exam/user-list", handler.postAddUser)
	e.GET("/exam/user-list/:id", handler.getUserById)
	e.PUT("/exam/user-list", handler.putUpdateUser)
	return handler
}

func (h *ExamHandler) getHelloWorld(c echo.Context) error {
	helloText := h.uc.ExamHelloWorldText()
	println(helloText)
	return c.JSON(http.StatusOK, helloText)
}

func (h *ExamHandler) getAlluser(c echo.Context) error {
	userList, err := h.uc.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if userList == nil {
		c.JSON(http.StatusNoContent, []*exam.ExamUser{})
	}
	ip := c.Request().Header.Get("X-Forwarded-For")
	print(ip)
	return c.JSON(http.StatusOK, userList)

}

func (h *ExamHandler) postAddUser(c echo.Context) error {
	name := c.Request().FormValue("name")
	email := c.Request().FormValue("email")

	if name == "" {
		return c.JSON(http.StatusBadRequest, "name field is required !")
	}
	if email == "" {
		return c.JSON(http.StatusBadRequest, "email field is required !")
	}
	err := h.uc.PostAdduser(name, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "status ok",
	})
}

func (h *ExamHandler) getUserById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userDetail, err := h.uc.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if userDetail == nil {
		// Note ::  StatusNoContent [204] can't return any json if want to return error must change to other status ex.404
		return c.JSON(http.StatusNoContent, map[string]string{})
		// Alternative when want to return empty json
		// return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, userDetail)
}

func (h *ExamHandler) putUpdateUser(c echo.Context) error {
	idParam := c.Request().FormValue("id")
	name := c.Request().FormValue("name")
	email := c.Request().FormValue("email")
	id, err := strconv.Atoi(idParam)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if name == "" {
		return c.JSON(http.StatusBadRequest, "name field is required !")
	}
	if email == "" {
		return c.JSON(http.StatusBadRequest, "email field is required !")
	}

	updateDetail, err := h.uc.PutUpdateUser(id, name, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updateDetail)
}
