package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Path: ")
	var path string
	path, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("incorrect path: %s", err)
	}

	/*
		This is for windows
	*/
	path = strings.TrimSpace(path)
	path = strings.TrimSuffix(path, "\n")
	path = strings.TrimSuffix(path, "\r")

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "error:") {
			fmt.Println(text)
		}
	}

	file.Close()
}
