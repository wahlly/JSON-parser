package utils

import "encoding/json"

func ValidateJson(data string) bool {
	var js map[string] interface {}
	//parse the json encoded data, and store in js
	err := json.Unmarshal([]byte(data), &js)
	return err == nil
}