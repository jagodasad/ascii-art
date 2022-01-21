package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

/* You’d need to run this code on your local machine */
/* Because, play.golang.org doesn't allow files */

func main() {
	art() // run the ANSI art function

	if err := keyboard.Open(); err != nil { // quick keypress test/loop
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
		}
		if string(char) == "r" || string(char) == "R" { // reload
			art()
		}
		if string(char) == "q" || string(char) == "Q" || key == keyboard.KeyEsc { // exit
			os.Exit(0)
		}
	}
}

func art() {
	// Example ANSI art file: https://16colo.rs/pack/mist0121/raw/LDA-PHASE90.ANS
	writeCp437("/home/robbiew/go/src/github.com/robbiew/ansi-col-count/LDA-PHASE90.ANS", 80) // 80 is the max line length
}

func writeCp437(file string, limit int) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	noSauce := TrimStringFromSauce(string(content)) // strip off the SAUCE metadata
	s := bufio.NewScanner(strings.NewReader(string(noSauce)))

	for s.Scan() {

		var d io.Reader = strings.NewReader(s.Text())                    // each line as string
		utf8 := transform.NewReader(d, charmap.CodePage437.NewDecoder()) // decode from CP437 to UTF-8
		decBytes, _ := ioutil.ReadAll(utf8)
		decS := string(decBytes)

		f := wrap.NewWriter(limit) // reflow packag from github.com/muesli/reflow
		f.PreserveSpace = true
		f.Newline = []rune{'\n'}
		f.KeepNewlines = true
		f.Write([]byte(decS))

		var cp437 io.Reader = strings.NewReader(f.String())
		cp437 = transform.NewReader(cp437, charmap.CodePage437.NewEncoder()) // encode bytes to CP437
		encBytes, _ := ioutil.ReadAll(cp437)
		encB := string(encBytes)
		fmt.Println(encB)
		time.Sleep(70 * time.Millisecond) // wait for a bit between lines
	}
}

func TrimStringFromSauce(s string) string {
	if idx := strings.Index(s, "COMNT"); idx != -1 {
		string := s
		delimiter := "COMNT"
		leftOfDelimiter := strings.Split(string, delimiter)[0]
		trim := TrimLastChar(leftOfDelimiter)
		return trim
	}
	if idx := strings.Index(s, "SAUCE00"); idx != -1 {
		string := s
		delimiter := "SAUCE00"
		leftOfDelimiter := strings.Split(string, delimiter)[0]
		trim := TrimLastChar(leftOfDelimiter)
		return trim
	}
	return s
}

func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
