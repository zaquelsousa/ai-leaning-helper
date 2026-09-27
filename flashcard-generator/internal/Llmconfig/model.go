package llmconfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type LLM struct{
	ModeUrl string
	Model string
	TokenIn int
	TokenOut int
	Duration int64
}


type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`

}

type ChatResponse struct {
	Model   string  `json:"model"`
	Message Message `json:"message"`
	Done    bool    `json:"done"`
	PromptEvalCount    int     `json:"prompt_eval_count"`
	EvalCount          int     `json:"eval_count"`
	PromptEvalDuration int64   `json:"prompt_eval_duration"`
	EvalDuration       int64   `json:"eval_duration"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}


func (m *LLM)SetModel(){
	m.ModeUrl = "http://localhost:11434/api/chat"
	m.Model = "qwen3:4b"
}

func (a *LLM)Chat(notesPath string)(ChatResponse, error){
	studyNote, err := os.ReadFile(notesPath)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	text := string(studyNote)

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
		Model: a.Model,
		Messages: []Message{msg},
		Stream: false,
	}
	msgInJson, err := json.Marshal(message)
	
	if err != nil {
		return ChatResponse{}, err
	}

	reqBody := bytes.NewBuffer(msgInJson)
	response, err := http.Post("http://localhost:11434/api/chat", "application/json", reqBody)
	
	if err != nil {
		return ChatResponse{}, err
	}
	defer response.Body.Close()

	var result ChatResponse

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return ChatResponse{}, err
	}
	
	a.TokenIn += result.PromptEvalCount
	a.TokenOut += result.EvalCount
	a.Duration += result.EvalDuration

	return result, nil
}
