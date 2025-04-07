package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("BookTracker - Track your reading journey")
	books := Books{}
	storage := NewStorage[Books]("books.json")
	if err := storage.Load(&books); err != nil {
		fmt.Printf("Error loading json")
		os.Exit(1)
	}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&books)
	if err := storage.Save(books); err != nil {
		fmt.Printf("Error save json")
		os.Exit(1)
	}
}
