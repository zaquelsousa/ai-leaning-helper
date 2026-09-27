package scanner

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"os"
)



func OpenDatabase() (*sql.DB, error) {
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


func CalculateHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}



func HasDocumentChanged(db *sql.DB, path string, currentHash string) (bool, error){
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


func UpdateDocumentHash(db *sql.DB, path string, currentHash string) error {
	_, err := db.Exec(`
		INSERT INTO documents (path, content_hash, processed_at)
		VALUES (?, ?, datetime('now'))
		ON CONFLICT(path) DO UPDATE SET
			content_hash = excluded.content_hash,
			processed_at = datetime('now')
	`, path, currentHash)

	return err
}

