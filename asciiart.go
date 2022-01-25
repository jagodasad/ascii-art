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
	wordOneRune := []rune(wordArg[1])
	/*newLine := ""
	for _, i := range wordArg[1] {
		if i == '\\' {
			newLine = "\n"
		}
	}*/

	wordString := ""

	for i := 0; i < len(wordArg[1]); i++ {
		if wordOneRune[i] != '\\' {
			wordString = wordString + string(wordOneRune[i])
		} else if i != 0 && i != len(wordArg[1])-1 && wordOneRune[i] == '\\' {
			wordString = wordString + " "
			wordString = wordString + string(wordOneRune[i])
			wordString = wordString + string(wordOneRune[i+1])
			wordString = wordString + " "
			i = i + 1
		}
	}

	// fmt.Println(wordString)

	// split := strings.Split(wordArg[1], "\\n")
	// splitString := strings.Join(split, "")

	result := strings.Fields(wordString)

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
		} else if result[a] == "\\n" {
			fmt.Print("\n")
		}
	}
}
