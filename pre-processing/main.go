package main

import (
	"bufio"
	"log"
	"os"
)

func scan_stdin() error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

	return scanner.Err()
}


func main() {
	err := scan_stdin()
	if err != nil {
		log.Fatal(err)
	}
}