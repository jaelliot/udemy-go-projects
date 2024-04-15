package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Todo saved successfully")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Note saved successfully")

}

func saveData() {}

func getNoteData() (string, string) {

	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
