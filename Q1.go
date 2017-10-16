package main

import (
  "html/template"
  "log"
  "net/http"
)

type PageVariables struct {
	Title string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

    title := "Guessing Game"
    HomePageVars := PageVariables{ //store the title in a struct
      Title : title,
    }

    t, err := template.ParseFiles("Q1.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}