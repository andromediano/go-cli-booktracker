package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add      string
	Del      int
	Edit     string
	MarkRead int
	Rate     string
	List     bool
	Command  string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	// 하이픈 옵션도 기존 코드와 호환성을 위해 남겨두지만 사용되지 않음
	flag.StringVar(&cf.Add, "add", "", "Add a new book. Format: 'Title:Author'")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a book by index. Format: 'index:Title:Author' (either Title or Author can be empty)")
	flag.IntVar(&cf.Del, "del", -1, "Delete a book by index")
	flag.IntVar(&cf.MarkRead, "read", -1, "Mark a book as read/unread by index")
	flag.StringVar(&cf.Rate, "rate", "", "Rate a book. Format: 'index:rating(1-5)'")
	flag.BoolVar(&cf.List, "list", false, "List all books")

	flag.Parse()

	// 하이픈이 없는 명령어 처리
	args := flag.Args()
	if len(args) > 0 {
		cf.Command = args[0] // 첫 번째 인자는 명령어 (add, list, read 등)
	}

	return &cf
}

func (cf *CmdFlags) Execute(books *Books) {
	// 하이픈 명령어 처리 (기존 플래그로 파싱된 경우)
	if cf.List || cf.Add != "" || cf.Edit != "" || cf.MarkRead != -1 || cf.Rate != "" || cf.Del != -1 {
		cf.executeWithFlags(books)
		return
	}

	// 하이픈 없는 명령어 처리
	if cf.Command != "" {
		cf.executeWithCommand(books)
		return
	}

	// 명령어가 없는 경우 목록 표시
	books.print()
}

func (cf *CmdFlags) executeWithFlags(books *Books) {
	switch {
	case cf.List:
		books.print()
	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for add. Please use Title:Author")
			os.Exit(1)
		}
		books.add(parts[0], parts[1])
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 3)
		if len(parts) != 3 {
			fmt.Println("Error, invalid format for edit. Please use index:Title:Author")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		// 함수의 반환 오류값을 확인하지 않으면 아래 메세지가 보여질것이다.
		// "errcheck: Error return value of books.edit is not checked"
		if err := books.edit(index, parts[1], parts[2]); err != nil {
			fmt.Printf("Error editing book: %v\n", err)
			os.Exit(1)
		}

	case cf.MarkRead != -1:
		if err := books.markAsRead(cf.MarkRead); err != nil {
			fmt.Printf("Error marking book as read: %v\n", err)
			os.Exit(1)
		}

	case cf.Rate != "":
		parts := strings.SplitN(cf.Rate, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for rating. Please use index:rating(1-5)")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for rating")
			os.Exit(1)
		}

		rating, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error: invalid rating value")
			os.Exit(1)
		}

		if err := books.rate(index, rating); err != nil {
			fmt.Printf("Error rating book: %v\n", err)
			os.Exit(1)
		}

	case cf.Del != -1:
		if err := books.delete(cf.Del); err != nil {
			fmt.Printf("Error deleting book: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid command")
	}
}

func (cf *CmdFlags) executeWithCommand(books *Books) {
	args := flag.Args()
	command := args[0]

	switch command {
	case "list":
		books.print()
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: add command requires a book title and author. Format: add 'Title:Author'")
			os.Exit(1)
		}
		parts := strings.SplitN(args[1], ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for add. Please use Title:Author")
			os.Exit(1)
		}
		books.add(parts[0], parts[1])
	case "edit":
		if len(args) < 2 {
			fmt.Println("Error: edit command requires parameters. Format: edit 'index:Title:Author'")
			os.Exit(1)
		}
		parts := strings.SplitN(args[1], ":", 3)
		if len(parts) != 3 {
			fmt.Println("Error, invalid format for edit. Please use index:Title:Author")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		if err := books.edit(index, parts[1], parts[2]); err != nil {
			fmt.Printf("Error editing book: %v\n", err)
			os.Exit(1)
		}
	case "read":
		if len(args) < 2 {
			fmt.Println("Error: read command requires a book index")
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid index for read command")
			os.Exit(1)
		}
		if err := books.markAsRead(index); err != nil {
			fmt.Printf("Error marking book as read: %v\n", err)
			os.Exit(1)
		}
	case "rate":
		if len(args) < 2 {
			fmt.Println("Error: rate command requires parameters. Format: rate 'index:rating'")
			os.Exit(1)
		}
		parts := strings.SplitN(args[1], ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for rating. Please use index:rating(1-5)")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for rating")
			os.Exit(1)
		}

		rating, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error: invalid rating value")
			os.Exit(1)
		}

		if err := books.rate(index, rating); err != nil {
			fmt.Printf("Error rating book : %v\n", err)
			os.Exit(1)
		}
	case "del":
		if len(args) < 2 {
			fmt.Println("Error: del command requires a book index")
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid index for delete")
			os.Exit(1)
		}
		if err := books.delete(index); err != nil {
			fmt.Printf("Error deleting book : %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command. Available commands: list, add, edit, read, rate, del")
	}
}
