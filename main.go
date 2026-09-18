package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file>")
		return
	}

	filePath := os.Args[1]

	studyNote, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
 	
	text := string(studyNote)

	//fmt.Println(text)

 	
	msg := Message{
		Role: "user",
		Content: `Read the following study notes and generate 5 flashcards. 
		Rules:
		- Questions must test understanding, not simple word matching.
		- Answers must be concise and grounded in the notes.
		- Do not invent information.
		- Return ONLY a valid JSON array.
		- Each object must contain "question" and "answer".

		Study notes:` + "\n" + text,
	}

	message := ChatRequest{
		Model: "qwen3:4b",
		Messages: []Message{msg},
		Stream: false,
	}
	
	
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		panic("Failed to marshal JSON: " + err.Error())
	}

	reqBody := bytes.NewBuffer(jsonBytes)

	start := time.Now()
	response, err := http.Post("http://localhost:11434/api/chat", "application/json", reqBody)
	elapsed := time.Since(start)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: \n", response.StatusCode)

	var result ChatResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	var flashcards []Flashcard

	err = json.Unmarshal([]byte(result.Message.Content), &flashcards)
	if err != nil{
		fmt.Println("Error parsing flashcards:", err)
		fmt.Println("Raw response:", result.Message.Content)
		return
	}

	for i, card := range flashcards {
		fmt.Printf("%d. %s\n", i+1, card.Question)
		fmt.Printf("Answer: %s\n\n", card.Answer)
	}


	//fmt.Println(result.Message.Content)
	fmt.Printf("\nresponse time: %s\n", elapsed)
	
}
