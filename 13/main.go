// ? 13 Continuing the above program modify the route to receive
// ? a querystring parameter called name and print the
// ? received name in "Hello <name>". Upon opening the
// ? URL http://localhost:8090/hello?name=Bala it should
// ? display Hello Bala

package main

import (
	"fmt"
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
	query := r.URL.Query()
	name := query.Get("name")

	if name != "" {
		fmt.Fprintf(w, "<h1>hello %v</h1>", name)
		return
	}

	w.Write([]byte("<h1>hello world</h1>"))
}
