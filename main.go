package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/wahlly/JSON-parser/utils"
)

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