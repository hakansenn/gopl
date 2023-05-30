package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler) // each request calls handler
	http.HandleFunc("/deneme",handler2) // each request calls handler

	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n",r.URL.Path)
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "deneme")
}