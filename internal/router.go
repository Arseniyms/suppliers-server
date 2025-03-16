package router

import (
	"arseniyms/suppliers/server/internal/controllers/auth"
	"arseniyms/suppliers/server/internal/controllers/companies"
	"arseniyms/suppliers/server/internal/controllers/password"
	"arseniyms/suppliers/server/internal/external/services/mail"
	"context"
	"fmt"
	"log"
	"net/http"
)

func HandleResponse() {
	client := companies.ConnectToDB()
	password.SetClient(client)
	defer client.Disconnect(context.Background())

	http.HandleFunc("/", companies.GetSuccess)
	http.HandleFunc(companies.COMPANIES_PATH, handleCompanies)
	http.HandleFunc(auth.LOGIN_PATH, handeLogin)
	http.HandleFunc(mail.MAIL_PATH, handleMail)
	http.HandleFunc("/import", handleImport)
	http.HandleFunc(password.PASSWORD_PATH, handlePassword)

	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	switch r.Method {
	case http.MethodGet:
		companies.GetCompanies(w, r)
	case http.MethodDelete:
		companies.DeleteCompany(w, r)
	case http.MethodPost:
		companies.CreateCompany(w, r)
	case http.MethodPatch:
		companies.PatchCompany(w, r)
	case http.MethodOptions:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handeLogin(w http.ResponseWriter, r *http.Request) {

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
}

func handleMail(w http.ResponseWriter, r *http.Request) {
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
}

func handleImport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodPost:
		companies.CreateCompanyArray(w, r)
	case http.MethodOptions:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		password.IsPasswordEmpty(w, r)
	case http.MethodPost:
		password.CreatePassword(w, r)
	case http.MethodDelete:
		password.DeletePassword(w, r)
	case http.MethodOptions:
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
