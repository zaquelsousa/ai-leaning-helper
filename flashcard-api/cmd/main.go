package main

import (
	"fmt"
	"log"
	"net/http"
)

type Flashcard struct{
	Question string `json:"question"`
	Answer string `json:"answer"`
}


func GetFlashCard(w http.ResponseWriter, r *http.Request){
	
}


func main(){
	fcard := Flashcard{
		Question: "quem descobriu o brasil",
		Answer: "tadin ele ta com frio agr",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, fcard)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
