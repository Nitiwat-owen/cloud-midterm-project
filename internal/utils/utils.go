package utils

import (
	"bufio"
	"os"
)

func GetFileContent(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		content = content + line
	}
	return content
}
