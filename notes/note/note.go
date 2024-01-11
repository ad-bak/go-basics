package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// here we use json tags to specify the names of the fields in the json file.
// the fields after `json:` will only be used when marshalling and unmarshalling
// in go marshalling is the process of converting a struct to a json string
// in javascript in can be compared to JSON.stringify()
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// here we don't need to use a pointer, but it's used just for practice
// we can add methods to structs by using the following syntax:
// func (receiver *StructName) MethodName() {}
// then we will be able to call the method on a struct instance like so:
// structInstance.MethodName()
// in this case n.Display()
func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n", n.Title, n.Content)
}

// here we use the os package to write the json to a file
// the 0644 is the file permissions, it's the same as the permissions you see when you run `ls -l`
func (n *Note) Save() error {
	// here we replace all spaces with underscores and make the title lowercase
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// here we marshal the note to json
	// as was said before, marshalling is the process of converting a struct to a json string
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("Invalid input.")
	}

	// here we use capital letters for the fields because we want them to be exported
	// if we used lowercase letters, the fields would be private and we wouldn't be able to access them
	// hence it wouldn't be saved in the json file
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
