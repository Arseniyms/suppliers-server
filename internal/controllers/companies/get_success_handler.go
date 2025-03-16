package companies

import "net/http"

func GetSuccess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully connected\nNow you can make requests"))
}
