package utils

func ValidateUser(username, password string) (string, bool) {
	if val, ok := users[username]; ok {
		if val == password {
			return "", true
		}
	}

	return "invalid username or password", false
}

var users map[string]string = map[string]string{
	"admin": "password",
	"editor": "secret",
	"trainer": "rahasia",
}