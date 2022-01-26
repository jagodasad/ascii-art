package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	wordArg := os.Args

	result := strings.Split(wordArg[1], "\\n")

	if wordArg[1] == "\\n" {
		result = append(result[:1])
	}

	f, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for a := 0; a < len(result); a++ {
		wordRune := []rune(result[a])
		if len(wordRune) != 0 && result[a] != "\\n" {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(wordRune); j++ {
					if lines[int(wordRune[j])*9-287+i] == "        " {
						fmt.Printf("        ")
					} else {
						fmt.Printf(lines[int(wordRune[j])*9-287+i])
					}
				}
				fmt.Print("\n")
			}
		} else {
			fmt.Print("\n")
		}
	}
}
