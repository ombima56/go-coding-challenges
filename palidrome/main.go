package main

import (
	"bufio"
	"fmt"
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

		if Palidrome(line) {
			fmt.Printf("%q is a palidrome.\n", line)
		} else {
			fmt.Printf("%q is not a palidrome.\n", line)
		}
	}
}
