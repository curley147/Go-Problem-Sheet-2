package main

import (
	"net/http"
	"text/template"
	"fmt"
	"math/rand"
	"strconv"
)
//struct to hold message 
type Messagedata struct{
	Message string
}

func guessHandler(w http.ResponseWriter, r *http.Request){
	//variable for number to be guessed
	var guessNum int
	//Adapted from https://github.com/data-representation/go-cookies/blob/master/go-cookie.go
	// Try to read the cookie.
	var cookie, err = r.Cookie("target")
	if err == nil {
		// If we could read it, try to convert its value to an int.
		cValInt, _ := strconv.Atoi(cookie.Value)
		guessNum = cValInt
	} 
	//if there's no cookie
	guessNum = rand.Intn(20)
	// Create a cookie instance and set the cookie.
	// You can delete the Expires line (and the time import) to make a session cookie.
	mycookie := &http.Cookie{
		Name:    "target",
		Value:   strconv.Itoa(guessNum),
	}
	http.SetCookie(w, mycookie)
	
	//first cookie attempt
	// c, err1 := r.Cookie("target")
	// if err1 != nil {
	// 	fmt.Println("cookie error", err1)
	// }
	// if c == nil || c.Name!="target" {
	// 	rand := rand.Intn(20)
	// 	randNumStr := strconv.Itoa(rand)
	// 	cookie := http.Cookie{Name: "target", Value: randNumStr}
	// 	http.SetCookie(w, &cookie)
	// }
	
	//parsing tmpl file using template package
	t, err1 := template.ParseFiles("template/guess.tmpl")
	if err1 != nil {
		//error check
		fmt.Println("template retrieval failed:", err1)
	}
	//sending Messagedata struct with response(value inject into {{.Message}} in tmpl file)
	t.Execute(w, Messagedata{Message: "Guess a number between 1 and 20"})
}
//main
func main() {
	//static file handler
	http.Handle("/", http.FileServer(http.Dir("./Ps2")))
	//template handler 
	http.HandleFunc("/guess", guessHandler)
	//Listen for requests on port 80
	http.ListenAndServe(":8080", nil)
}