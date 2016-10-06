package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

func removePunct(input chan string, output chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	re := regexp.MustCompile("[^a-zA-Z\\s]")
	for line := range input {
		output <- re.ReplaceAllString(line, "")
	}
	wg.Done()
}

func scanStdIn(output chan string, errors chan error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		output <- line
	}
	close(output)

	err := scanner.Err()

	if err == nil {
		return
	}

	errors <- err
}

func main() {

	stdInLines := make(chan string, 10000)
	stdInErrors := make(chan error)
	onlyAlpha := make(chan string, 10000)
	alphaWorkers := 10

	var wg sync.WaitGroup
	done := make(chan bool)

	stdoutLog := log.New(os.Stdout, "", 0)

	go scanStdIn(stdInLines, stdInErrors)

	go func() {
		for err := range stdInErrors {
			log.Fatal(err)
		}
	}()

	for n := 0; n < alphaWorkers; n++ {
		go removePunct(stdInLines, onlyAlpha, &wg)
	}

	go func() {
		wg.Wait()
		close(onlyAlpha)
	}()

	go func() {
		for line := range(onlyAlpha) {
			log.Println(line)
		}
		done <- true
	}()

	ticker := time.NewTicker(1 * time.Second)

RUN_LOOP:
	for {
		select {
			case <-ticker.C:
				stdoutLog.Printf("stdInLines len: %d, onlyAlpha len: %d", len(stdInLines), len(onlyAlpha))
			case <- done:
				break RUN_LOOP
		}
	}
}