package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
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

func toLower(input chan string, output chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for line := range input {
		output <- strings.ToLower(line)
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
	lowerCase := make(chan string, 10000)
	lowerWorkers := 10

	var wgAlpha, wgLower sync.WaitGroup
	done := make(chan bool)

	stdoutLog := log.New(os.Stdout, "", 0)

	go scanStdIn(stdInLines, stdInErrors)

	go func() {
		for err := range stdInErrors {
			log.Fatal(err)
		}
	}()

	for n := 0; n < alphaWorkers; n++ {
		go removePunct(stdInLines, onlyAlpha, &wgAlpha)
	}

	go func() {
		wgAlpha.Wait()
		close(onlyAlpha)
	}()

	for n := 0; n < lowerWorkers; n++ {
		go toLower(onlyAlpha, lowerCase, &wgLower)
	}

	go func() {
		wgLower.Wait()
		close(lowerCase)
	}()

	go func() {
		for line := range(lowerCase) {
			log.Println(line)
		}
		done <- true
	}()



	ticker := time.NewTicker(1 * time.Second)

RUN_LOOP:
	for {
		select {
			case <-ticker.C:
				stdoutLog.Printf("stdInLines len: %d, onlyAlpha len: %d, lowerCase len: %d", len(stdInLines), len(onlyAlpha), len(lowerCase))
			case <- done:
				break RUN_LOOP
		}
	}
}