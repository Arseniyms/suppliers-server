package auth

import "net/http"

func GetIsAuthentificated(r *http.Request) bool {
	token, err := ValidateConnection(r)

	return err == nil && token != ""
}
