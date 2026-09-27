package main

import (
	//"bytes"
	//"crypto/sha256"
	//"database/sql"
	//"encoding/hex"
	//"encoding/json"
	"fmt"
	//"io/fs"
	"learn-helper/internal/scanner"
	//"net"
	//"net/http"
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





func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file>")
		return
	}

	
	root := "/home/zakk/Desktop/computerScience/Operation-first-job/vaults/backend-roadmap"
	err := scanner.WalkDirRecursive(root)
	if err != nil{
		fmt.Println("error on scan notes")
		return
	}
	
	
	/*scan the designated dir for notes
	also we need blaclist notes
	//open the db
	db, err := openDatabase()
	if err != nil {
		fmt.Println("Database error:", err)
		return
	}
	defer db.Close()
	fmt.Println("Database opened successfully")

	//scan the dir for .md files

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
	*/


	 
	/*LLM call so it gens the cards
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
	}
	*/
	
	

	/* we send to the API so it can insert on db and see on frontend
	
	//in go _ mean that we dont care about the idx but we want the actual element
	for _, card := range flashcards { 
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
	
	*/


	//metrics
	//fmt.Printf("\nresponse time: %s\n", elapsed)
	
}
