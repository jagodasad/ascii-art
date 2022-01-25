package main

import (
	"bufio"
	"fmt"
	"strings"

	//"io/ioutil"
	"log"
	"os"
	//"strings"
)

// using Bufio

func main() {
	wordArg := os.Args
	wordRune := []rune(wordArg[1])

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

	fmt.Println(lines[0])

	/*rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}*/

	for i := 0; i < 8; i++ {
		lines := strings.Split(string(wordRune), "\\n")
		// f.Newline = []rune{'\n'}
		// for l, line := range lines {
		for j := 0; j < len(wordRune); j++ {
			if lines[int(wordRune[j])*9-287+i] == "        " {
				fmt.Print("        ")
			} else {
				fmt.Print(lines[int(wordRune[j])*9-287+i])
			}
		}
		fmt.Print("\n")
	}
}
