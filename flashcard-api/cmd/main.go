package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Flashcard struct{
	Question string `json:"question"`
	Answer string `json:"answer"`
}


var flashcards = []Flashcard{
	{Question: "quem descobriu o brasil?", Answer: "tadin ele ta com frio"},
	{Question: "qual tamanho do meu penis?", Answer: "grande"},
}

func getAllFlashcards()[]Flashcard{
	return flashcards
}

func GetFlashCard(w http.ResponseWriter, r *http.Request){
	f := getAllFlashcards()
	

	json.NewEncoder(w).Encode(f)
}

func CreateFlashcard(w http.ResponseWriter, r *http.Request){
	var f Flashcard

	err := json.NewDecoder(r.Body).Decode(&f)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	flashcards = append(flashcards, f)
}


func corsMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "http://0.0.0.0:8081")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") 
        w.Header().Set("Access-Control-Allow-Credentials", "true") 

		if(r.Method == "OPTIONS"){
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main(){

	http.HandleFunc("/", GetFlashCard)
	http.HandleFunc("/flashcards", CreateFlashcard)


	handleCors := corsMiddleware(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", handleCors))
}
