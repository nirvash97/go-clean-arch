package pathList

func PublicPathConst() []string {
	publicPath := []string{
		"/auth/signUp",
		"/auth/signIn",
		"/exam/hello",
		"/exam/user-list",
		"/exam/user-list/:id",
	}
	return publicPath
}
