package fileread

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Read(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to read file")
		fmt.Print(err)
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("failed to read file")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read file")
	}
	file.Close()
	return lines, nil
}
func Write(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create file")
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Faild to convert data to JSON.")
	}
	file.Close()
	return nil
}
