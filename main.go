package main

import "fmt"

func main() {
  fmt.Println("BookTracker - Track your reading journey")
  books := Books{}
  storage := NewStorage[Books]("books.json")
  storage.Load(&books)
  cmdFlags := NewCmdFlags()
  cmdFlags.Execute(&books)
  storage.Save(books)
}
