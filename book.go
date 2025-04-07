package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Book struct {
	Title      string
	Author     string
	Read       bool
	AddedAt    time.Time
	FinishedAt *time.Time
	Rating     int // 1-5 star rating
}

type Books []Book

func (books *Books) add(title string, author string) {
	book := Book{
		Title:      title,
		Author:     author,
		Read:       false,
		FinishedAt: nil,
		AddedAt:    time.Now(),
		Rating:     0,
	}

	*books = append(*books, book)
}

func (books *Books) validateIndex(index int) error {
	if index < 0 || index >= len(*books) {
		err := errors.New("invalid book index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (books *Books) delete(index int) error {
	b := *books

	if err := b.validateIndex(index); err != nil {
		return err
	}

	*books = slices.Delete(b, index, index+1)

	return nil
}

func (books *Books) markAsRead(index int) error {
	b := *books

	if err := b.validateIndex(index); err != nil {
		return err
	}

	isRead := b[index].Read

	if !isRead {
		finishTime := time.Now()
		b[index].FinishedAt = &finishTime
	} else {
		b[index].FinishedAt = nil
	}

	b[index].Read = !isRead

	return nil
}

func (books *Books) rate(index int, rating int) error {
	b := *books

	if err := b.validateIndex(index); err != nil {
		return err
	}

	if rating < 1 || rating > 5 {
		err := errors.New("rating must be between 1 and 5")
		fmt.Println(err)
		return err
	}

	b[index].Rating = rating

	return nil
}

func (books *Books) edit(index int, title string, author string) error {
	b := *books

	if err := b.validateIndex(index); err != nil {
		return err
	}

	if title != "" {
		b[index].Title = title
	}

	if author != "" {
		b[index].Author = author
	}

	return nil
}

func (books *Books) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Author", "Status", "Added At", "Finished At", "Rating")
	for index, b := range *books {
		status := "Unread"
		finishedAt := ""
		rating := ""

		if b.Read {
			status = "Read"
			if b.FinishedAt != nil {
				finishedAt = b.FinishedAt.Format(time.RFC1123)
			}
		}

		if b.Rating > 0 {
			rating := ""
			for range b.Rating {
				rating += "⭐"
			}
		}

		table.AddRow(strconv.Itoa(index), b.Title, b.Author, status, b.AddedAt.Format(time.RFC1123), finishedAt, rating)
	}

	table.Render()
}
