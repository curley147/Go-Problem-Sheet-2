package main

import (
	"net/http"
	"text/template"
	"fmt"
)

type Messagedata struct{
	Message string
}

func guessHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("template/guess.tmpl")
	if err != nil {
		fmt.Println("template retrieval failed:", err)
	}
	t.Execute(w, Messagedata{Message: "Guess a number between 1 and 20"})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Ps2")))
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}