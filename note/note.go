package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	Createdon time.Time
}
func (n Note) Display() {
	fmt.Printf("your title is :%v \n and the content written inside it :%v", n.Title, n.Content)
}
func (n Note) Save() error {
	filename := strings.ReplaceAll(n.Title, "", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}
func New(title, content string) (Note, error) {
	if title == " " || content == "" {
		return Note{}, errors.New("you have not enter anything")
	}
	return Note{
		Title:     title,
		Content:   content,
		Createdon: time.Now(),
	}, nil
	}
