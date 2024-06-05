// ? 12. Create a route GET /hello - and return hello world  Upon opening the URL
// ?	 http://localhost:8090/hello Hello world should be displayed in browser

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handleGet)

	log.Fatal(http.ListenAndServe(":8090", r))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))
}
