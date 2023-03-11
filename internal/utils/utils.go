package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileContent(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	// Create a buffered reader for the file
	reader := bufio.NewReader(file)

	// Use asynchronous I/O to read the file contents
	buffer := make([]byte, 6*1024*1024) // 6 MB
	_, err := reader.Read(buffer)
	if err != nil {
			fmt.Println(err)
			return ""
	}

	// Print the file contents
	content := string(buffer)
	return content
}
