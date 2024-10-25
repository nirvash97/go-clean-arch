package http

import (
	"go-clean-arch/modules/entities/auth"
	"go-clean-arch/modules/usecase"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func (h *AuthHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte("one_wish")
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "username and password is required to operate", http.StatusBadRequest)
		return
	}
	// hashpassword, hashErr := h.uc.HashPassword(password)
	// if hashErr != nil {
	// 	http.Error(w, hashErr.Error(), http.StatusInternalServerError)
	// 	return
	// }
	userDetail, errDetail := h.uc.HandleSignIn(username)
	if errDetail != nil {
		http.Error(w, errDetail.Error(), http.StatusInternalServerError)
		return
	}
	// ComparehashAndPassword must compare between hashed password and plain text password
	err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(password))
	isPasswordValid := err == nil
	if isPasswordValid {
		expireTime := time.Now().Add(5 * time.Minute)
		claims := auth.Claims{
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expireTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(tokenString))
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
