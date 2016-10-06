package main

import (
	"bufio"
	"log"
	"os"
)

func scanStdIn(output chan string, errors chan error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		output <- line
	}

	err := scanner.Err()

	if err == nil {
		return
	}

	errors <- err
}


func main() {

	stdInLines := make(chan string)
	stdInErrors := make(chan error)

	go scanStdIn(stdInLines, stdInErrors)

	go func() {
		for err := range stdInErrors {
			log.Fatal(err)
		}
	}()

	for line := range(stdInLines) {
		log.Println(line)
	}
}