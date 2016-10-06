package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func removePunct(input chan string, output chan string) {
	re := regexp.MustCompile("[^a-zA-Z\\s]")
	for line := range input {
		output <- re.ReplaceAllString(line, "")
	}
}

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
	onlyAlpha := make(chan string)
	alphaWorkers := 10

	go scanStdIn(stdInLines, stdInErrors)

	go func() {
		for err := range stdInErrors {
			log.Fatal(err)
		}
	}()

	for n := 0; n < alphaWorkers; n++ {
		go removePunct(stdInLines, onlyAlpha)
	}

	for line := range(onlyAlpha) {
		log.Println(line)
	}
}