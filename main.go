package main

import (
	"fmt"
	"log"
	"net/http"

	"os/exec"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		fmt.Fprintf(w, string("Url Param 'key' is missing"))
		return
	}

	key := keys[0]
	description, err := exec.Command("go", "run", "cmd/godescribe/main.go", key).Output()

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(description))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
