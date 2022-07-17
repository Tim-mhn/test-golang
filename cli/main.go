package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "logs.txt", "path of the logs file")
	level := flag.String("level", "ERROR", "Log level: ERROR | INFO")

	flag.Parse()

	file, err := os.Open(*path)

	if err != nil {
		log.Fatal("File could not be opened")
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
