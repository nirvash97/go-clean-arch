package pathList

import "strings"

func IsPathPublic(path string) bool {
	prefixList := []string{
		"/exam/",
	}
	for _, prefix := range prefixList {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	var publicPath = map[string]struct{}{
		"/auth/signUp": {},
		"/auth/signIn": {},
	}
	_, isFound := publicPath[path]
	return isFound
}
