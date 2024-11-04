package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

func ValidateJson(data string) bool {
	var js map[string] interface {}
	//parse the json encoded data, and store in js
	err := json.Unmarshal([]byte(data), &js)
	return err == nil
}

func main() {
	fmt.Println("Please enter the path to the file")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)
	// fmt.Println(ValidateJson(val))
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if ValidateJson(string(fileContent)) {
		fmt.Println("JSON file is valid")
	} else{
		fmt.Println("JSON file is invalid")
	}
}