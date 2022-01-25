package main

import (
	"bufio"
	"fmt"

	//"strings"
	//"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	wordArg := os.Args
	newLine := ""
	for _, i := range wordArg[1] {
		if i == '\\' {
			newLine = "\n"
		}
	}

	split := strings.Split(wordArg[1], "\\n")
	splitString := strings.Join(split, "")

	wordRune := []rune(splitString)

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

	if len(wordRune) != 0 {
		for a := 0; a < len(split); a++ {
			fmt.Print("\n")
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
		}
	}
	if newLine != "" {
		fmt.Print(newLine)
	}
}
