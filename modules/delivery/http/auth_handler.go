package http

import (
	"go-clean-arch/modules/entities/auth"
	"go-clean-arch/modules/usecase"
	"net/http"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "`username` field is required !", http.StatusBadRequest)
		return
	}
	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "`password` field is required !", http.StatusBadRequest)
		return
	}

	mail := r.FormValue("mail")
	if mail == "" {
		http.Error(w, "`mail` field is required !", http.StatusBadRequest)
		return
	}

	isExist := h.uc.IsUsernameExist(username)
	if isExist {
		http.Error(w, username+" already exist", http.StatusBadRequest)
		return
	}
	hashpassword, hashErr := h.uc.HashPassword(password)
	if hashErr != nil {
		http.Error(w, hashErr.Error(), http.StatusInternalServerError)
		return
	}
	signUpDetail := auth.UserAuth{
		Username: username,
		Password: hashpassword,
		Email:    mail,
	}

	err := h.uc.HandleSignUp(signUpDetail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	//decodeData := json.Encoder(map[string]string{"err" :""})

	//decodeJson := json.NewDecoder(Body).Decode(&signUpDetail)

}
