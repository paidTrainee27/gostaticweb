package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)

	fmt.Printf("server running on :%d \n", port)
	fmt.Println("press ctrl+c to terminate")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error Parsing form ", http.StatusBadRequest)
	}
	name := r.FormValue("name")
	add := r.FormValue("address")
	fmt.Fprintf(w, "Your name: %s \nand address: %s", name, add)
}
