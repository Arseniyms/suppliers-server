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
	http.HandleFunc("/users", connectors.GetData)
	http.HandleFunc("/users/add", connectors.CreateItem)
	http.HandleFunc("/users/patch", connectors.PatchItem)

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			connectors.DeleteItem(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// http.ListenAndServe("0.0.0.0:9090", nil)

	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
