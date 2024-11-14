package auth

type UserAuth struct {
	UserId   string `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
