package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	//"io/fs"
	"net/http"
	"os"
	//"path/filepath"
	//"strings"
	//"time"

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

func calculateHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}


func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "documents.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS documents (
			id INTEGER PRIMARY KEY,
			path TEXT NOT NULL UNIQUE,
			content_hash TEXT NOT NULL,
			processed_at TEXT NOT NULL
		);
	`)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func hasDocumentChanged(db *sql.DB, path string, currentHash string) (bool, error){
	var storeHash string

	err := db.QueryRow(`
		SELECT content_hash
		FROM documents
		WHERE path = ?
	`, path).Scan(&storeHash)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	return storeHash != currentHash, nil
}

func updateDocumentHash(db *sql.DB, path string, currentHash string) error {
	_, err := db.Exec(`
		INSERT INTO documents (path, content_hash, processed_at)
		VALUES (?, ?, datetime('now'))
		ON CONFLICT(path) DO UPDATE SET
			content_hash = excluded.content_hash,
			processed_at = datetime('now')
	`, path, currentHash)

	return err
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file>")
		return
	}

	/*
	//open the db
	db, err := openDatabase()
	if err != nil {
		fmt.Println("Database error:", err)
		return
	}
	defer db.Close()
	fmt.Println("Database opened successfully")

	//scan the dir for .md files
	root := "/home/zakk/Desktop/computerScience/Operation-first-job/vaults/backend-roadmap"

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(d.Name(), ".md"){
			hash, err := calculateHash(path)
			
			if err != nil {
				return err
			}
			
			changed, err := hasDocumentChanged(db, path, hash)
			
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			if !changed {
				fmt.Println("Skip:", path)
				return nil
			}

			fmt.Println("Process: ", path)

			err = updateDocumentHash(db, path, hash)
			if err != nil {
				return err
			}


			//fmt.Printf("File: %s\n", path)
			//fmt.Printf("Hash: %s\n\n", hash)
		}

		return nil
	})
	
	if err != nil {
		fmt.Println("Error:", err)
	}

	
	//parse a md file so i can use with the prompt
	filePath := os.Args[1]
	studyNote, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error: ", err)
		return
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
	}*/
	
	flashcards := []Flashcard{
		{
			Question: "What is HTTP?",
			Answer:   "A protocol used for communication between clients and servers.",
		},
		{
			Question: "What does GET do in HTTP?",
			Answer:   "It requests a representation of a resource.",
		},
		{
			Question: "What does POST do in HTTP?",
			Answer:   "It submits data to the server, commonly to create a resource.",
		},
		{
			Question: "What status code means 'Not Found'?",
			Answer:   "404 Not Found.",
		},
		{
			Question: "What status code means a resource was created?",
			Answer:   "201 Created.",
		},
	}
	
	//in go _ mean that we dont care about the idx but we want the actual element
	for _, card := range flashcards { 
		data, err := json.Marshal(card)
		if err != nil {
			fmt.Printf("Error parsing cards")
		}
		
		resp, err := http.Post("http://localhost:8080/flashcards", "application/json", bytes.NewReader(data))
		if err != nil {
			fmt.Printf("error while createing the flashcard on the API")
		}
	
		fmt.Println(resp.Status)
		defer resp.Body.Close()

	}


	//fmt.Printf("\nresponse time: %s\n", elapsed)
	
}
