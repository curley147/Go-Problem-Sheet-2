package main

import (
  "net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Ps2")))
	http.ListenAndServe(":8080", nil)
}

// func HomePage(w http.ResponseWriter, r *http.Request){
    
// }