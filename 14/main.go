// ? 14. Create POST route called /postHello. The request body should
// ? 	be a json of the format { name: Bala}. Upon calling this API
// ? 	from POSTMAN it should return Hello Bala

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string
}

func handlePost(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if r.Body == nil || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": name not found}`))
		return
	}

	fmt.Fprintf(w, `{"message": "hello %v"}`, user.Name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/postHello", handlePost)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8090", r))
}
