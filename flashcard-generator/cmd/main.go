package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	llmconfig "learn-helper/internal/Llmconfig"
	"learn-helper/internal/scanner"
	"net/http"

	_ "modernc.org/sqlite"
)

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Model   string  `json:"model"`
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Flashcard struct {
    Question string `json:"question"`
    Answer   string `json:"answer"`
}





func main(){
	//root := "/home/zakk/Desktop/computerScience/Operation-first-job/vaults/backend-roadmap"
	test := "/home/zakk/test"

	notes, err := scanner.WalkDirRecursive(test)
	if err != nil{
		fmt.Println("error on scan notes")
		return
	}

	
	
	var model  llmconfig.LLM

	model.SetModel()

	var allFlashcards []Flashcard

	for _, note := range notes{
		var flashcards []Flashcard

		resp, err := model.Chat(note)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal([]byte(resp.Message.Content), &flashcards)
		if err != nil{
			fmt.Println("Error parsing flashcards:", err)
			fmt.Println("Raw response:", resp.Message.Content)
			return
		}
		
		allFlashcards = append(allFlashcards, flashcards...)
		//this ... means that go will apeand the element rather then the slices
	}

	
	//in go _ mean that we dont care about the idx but we want the actual element
	for _, card := range allFlashcards { 
		data, err := json.Marshal(card)
		if err != nil {
			fmt.Printf("Error parsing cards")
		}
		
		resp, err := http.Post("http://localhost:8080/flashcards", "application/json", bytes.NewReader(data))
		if err != nil {
			fmt.Printf("error while createing the flashcard on the API slk")
		}
	
		fmt.Println(resp.Status)
		defer resp.Body.Close()

	}
	
	//metrics
	fmt.Printf("Job finish with:")
	fmt.Println("Token in: ", model.TokenIn)
	fmt.Println("Token out: ", model.TokenOut)
	fmt.Println("Duration: ", model.Duration)
	
}
