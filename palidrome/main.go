package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Error: usage <go run.go> <data.txt>")
		return
	}

	arg := os.Args[1]

	file, err := os.Open(arg)
	if err != nil {
		log.Printf("Failed to open: %s file", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
	}
}