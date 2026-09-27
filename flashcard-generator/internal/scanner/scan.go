package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)


var blacklist = []string{
    "Docker-for-Backend-Development.md",
    "go-roadmap.md",
	"linux_backend_development_guide.md",
	"linux_daily_driver_guide.md",
	"networking_for_backend_development.md",
	"Operating-Systems-for-Backend-Engineering.md",
	"REST-APIs-for-Backend-Development.md",
	"testing_for_backend_development.md",
}

func isBlacklisted(path string, blacklist []string)bool{
	for _, file := range blacklist {
		if file == filepath.Base(path) {
			return true
		}
	}
	return false
}

func WalkDirRecursive(root string)([]string, error){
	var notes []string
	//creeate the sql file where the process will store the state of
	//the notes
	db, err := OpenDatabase()
	if err != nil{
		fmt.Println("Error while Open sql db")
	}
	defer db.Close()
	fmt.Println("Database opened successfully")


	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil{
			return err
		}

		if d.IsDir(){
			return nil
		}

		if strings.HasSuffix(d.Name(), ".md"){
			hash, err := CalculateHash(path)

			if err != nil {
				return err
			}

			if isBlacklisted(path, blacklist){
				return nil
			}

			changed, err := HasDocumentChanged(db, path, hash)

			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}

			if !changed {
				fmt.Println("Skip: ", path)
				return nil
			}
			
			fmt.Println("Process: ", path)

			err = UpdateDocumentHash(db, path, hash)
			notes = append(notes, path)

			if err != nil {
				return err
			}

			//fmt.Printf("File: %s\n", path)
			//fmt.Printf("Hash: %s\n\n", hash)
		}

		return nil
	})
	
	return notes, nil
}
