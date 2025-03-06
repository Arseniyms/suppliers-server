package main

import (
	"arseniyms/suppliers/server/connectors"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := connectors.ConnectToDB()
	defer client.Disconnect(context.Background())

	http.HandleFunc("/", connectors.GetSuccess)

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	// http.ListenAndServe("0.0.0.0:9090", nil)

	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
