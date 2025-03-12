package main

import (
	"arseniyms/suppliers/server/auth"
	"arseniyms/suppliers/server/connectors"
	"arseniyms/suppliers/server/mail"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := connectors.ConnectToDB()
	defer client.Disconnect(context.Background())

	http.HandleFunc("/", connectors.GetSuccess)

	http.HandleFunc(mail.MAIL_PATH, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		switch r.Method {
		case http.MethodPost:
			mail.SendToMail(w, r)
		case http.MethodOptions:
			return
		default:
			http.Error(w, "Mail Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc(auth.LOGIN_PATH, func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		switch r.Method {
		case http.MethodPost:
			auth.Login(w, r)
		case http.MethodGet:
			auth.ProtectedEndpoint(w, r)
		case http.MethodOptions:
			return
		default:
			http.Error(w, "Login Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc(connectors.COMPANIES_PATH, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		switch r.Method {
		case http.MethodGet:
			connectors.GetData(w, r)
		case http.MethodDelete:
			connectors.DeleteItem(w, r)
		case http.MethodPost:
			connectors.CreateItem(w, r)
		case http.MethodPatch:
			connectors.PatchItem(w, r)
		case http.MethodOptions:
			return
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/import", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case http.MethodPost:
			connectors.CreateArrayItems(w, r)
		case http.MethodOptions:
			return
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// http.ListenAndServe("0.0.0.0:9090", nil)

	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
