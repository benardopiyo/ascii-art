package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read from cmd-line
	strWord := os.Args[1]

	if strWord == "\\n" {
		fmt.Println()
		return

	} else if strWord == "" {
		return
	}

	words := strings.Split(strWord, "\\n")

	strArray := strings.Split(string(fileContent), "\n")


		for _, word := range words {
			if word == "" || word == "\n" {
				fmt.Println()
				continue
			}
			for i := 0; i < 8; i++ {
				for _, char := range word {
					startPoint := int(((char - 32) * 9) + 1)
					fmt.Print(strArray[startPoint+i])
				}
				fmt.Println()
			}

	}
}
