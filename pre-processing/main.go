package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}