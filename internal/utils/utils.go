package utils

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func GetFileContent(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewReader(file)

	// Use a bytes.Buffer to store the file content
	var content bytes.Buffer
	
	// Use a loop to read the file content until the end is reached
	for {
			buffer := make([]byte, 6*1024*1024) // 6 MB
			n, err := reader.Read(buffer)
			if err != nil {
					if err != io.EOF {
							return ""
					}
					break
			}
			content.Write(buffer[:n])
	}
	
	return content.String()
}
