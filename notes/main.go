package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

// In Go, interfaces define a set of methods to be implemented by a type.
// An interface, like 'saver', declares required methods, in this case, Save() returning an error.
// Types like Note and Todo implement the saver interface by defining Save() with the same signature.
// This is analogous to JavaScript's abstract classes: saver is the abstract class, Note and Todo are the inheriting classes.
// Both Note and Todo can be passed to functions like saveData() as they adhere to the saver interface,
// specifically by implementing the required Save() method.
// The key is matching the function type and return statement with the interface's method signature.

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()

	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

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
