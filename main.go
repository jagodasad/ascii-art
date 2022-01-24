// fmt.Println()
// file, err := os.Open("/path/to/file.txt") // opens file object
// if err != nil {
// 	log.Fatal(err)
// }
// defer file.Close()

// scanner := bufio.NewScanner(file) // scanner (could be 'lines') that scans line-by-line
// // optionally, resize scanner's capacity for lines over 64K, see next example
// for scanner.Scan() {
// 	fmt.Println(scanner.Text())
// }

// if err := scanner.Err(); err != nil { // very important it yields an error in case e.g. there was buffer full and
// 	log.Fatal(err) // all lines were not scanned through
// }

// ---------------------------
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := " #  ##   ## ##  ### ###  ## # # ###  ## # # #   # # ###  #  ##   #  ##   ## ### # # # # # # # # # # ### ### "
	s2 := "# # # # #   # # #   #   #   # #  #    # # # #   ### # # # # # # # # # # #    #  # # # # # # # # # #   #   # "
	s3 := "### ##  #   # # ##  ##  # # ###  #    # ##  #   ### # # # # ##  # # ##   #   #  # # # # ###  #   #   #   ## "
	s4 := "# # # # #   # # #   #   # # # #  #  # # # # #   # # # # # # #    ## # #   #  #  # # # # ### # #  #  #       "
	s5 := "# # ##   ## ##  ### #    ## # # ###  #  # # ### # # # #  #  #     # # # ##   #  ###  #  # # # #  #  ###  #  "

	lex := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "?"}

	var lines []string
	lines = append(lines, s1, s2, s3, s4, s5)

	dict := Dictionnary{}
	dict.Init(lex, lines, 4)

	// a := dict.Letters["A"]
	// a.Print()

	// f := dict.Letters["F"]
	// f.Print()

	// dict.Letters["A"].Print()
	// dict.Print()

	dict.PrintWord("Hello")
	// dict.PrintWord("MANHATTAN")
	// dict.PrintWord("M@NHATTAN")
	// dict.PrintWord("ManhAtTan")
	// dict.PrintWord("Lorem ipsum dolor sit amet,...")
}

type Dictionnary struct {
	Height, Length int
	Letters        map[string]Letter
}

func (d *Dictionnary) Init(e []string, l []string, x int) {
	d.Height = len(l)
	d.Length = x
	d.Letters = make(map[string]Letter)
	for _, v := range e {
		var letter Letter
		lines := map[int]string{}
		letter.Lines = lines
		d.Letters[v] = letter
	}
	for k, v := range l {
		n := splitStr(v, x, 0, len(l))
		for i, j := range n {
			d.Letters[e[i]].Lines[k] = j
		}
	}
}

func (d *Dictionnary) PrintWord(s string) {
	for _, c := range s {
		fmt.Printf("%+q => %#U\n", string(c), c)
	}
	fmt.Println()

	out := make(map[int]string)
	for _, c := range s {
		for i := 0; i < d.Height; i++ {
			if findStr(strings.ToUpper(string(c)), Keys(d.Letters)) {
				x := d.Letters[strings.ToUpper(string(c))]
				out[i] += x.Lines[i]
			} else if unicode.IsSpace(c) {
				out[i] += strings.Repeat(" ", d.Length)
			} else {
				x := d.Letters["?"]
				out[i] += x.Lines[i]
			}
		}
	}

	for i := 0; i < len(out); i++ {
		fmt.Println(out[i])
	}
}

func (d *Dictionnary) Print() {
	for d, e := range d.Letters {
		fmt.Printf("index: %s\n", d)
		e.Print()
	}
}

type Letter struct {
	Lines map[int]string
}

func (l *Letter) Print() {
	for i := 0; i < len(l.Lines); i++ {
		fmt.Println(l.Lines[i])
	}
}

func splitStr(s string, l int, b int, e int) (r []string) {
	i := b
	e = len(s)
	for i < e {
		j := i + l
		if j > e {
			j = e
		}
		r = append(r, string(s[i:j]))
		i = j
	}
	return r
}

func findStr(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Keys(m map[string]Letter) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// -------------------------------------------------------

// package main

// import (
//     "fmt"
//     "os"
//     "bufio"
//     "strconv"
//     "strings"
//     "unicode"
// )

// var lex = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "?"}

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)

//     scanner.Split(bufio.ScanLines)

//     var L int
//     scanner.Scan()
//     fmt.Sscan(scanner.Text(),&L)
//     fmt.Fprintln(os.Stderr, "Letter's Length " + strconv.Itoa(L))

//     var H int
//     scanner.Scan()
//     fmt.Sscan(scanner.Text(),&H)
//     fmt.Fprintln(os.Stderr, "Letter's Height " + strconv.Itoa(H))

//     var T string
//     scanner.Scan()
//     fmt.Sscan(scanner.Text(),&T)
//     fmt.Fprintln(os.Stderr, "input text: " + T)

//     var Rows []string
//     // we have 27 item of 4 space each: 108 characters
//     for scanner.Scan() {
//         row := scanner.Text()
//         fmt.Fprintln(os.Stderr, row)
//         Rows = append(Rows, row)
//     }

//     dict := Dictionnary{}
// 	dict.Init(lex, Rows, L)

//     dict.PrintWord(T)
// // 	dict.PrintWord("Lorem ipsum dolor sit amet,... ")
// }

// type Dictionnary struct {
// 	Height, Length int
// 	Letters map[string]Letter
// }

// func (d *Dictionnary) Init(e []string, l []string, x int) {
// 	d.Height = len(l)
// 	d.Length = x
// 	d.Letters = make(map[string]Letter)
// 	for _, v := range e {
// 		var letter Letter
// 		lines := map[int]string{}
// 		letter.Lines = lines
// 		d.Letters[v] = letter
// 	}
// 	for k, v := range l {
// 		n := splitStr(v, x, 0, len(l))
// 		for i, j := range n {
// 			d.Letters[e[i]].Lines[k] = j
// 		}
// 	}
// }

// func (d *Dictionnary) PrintWord(s string) {
// 	out := make(map[int]string)
// 	for _, c := range s {
// 	    for i := 0; i < d.Height; i++ {
// 	        if findStr(strings.ToUpper(string(c)), Keys(d.Letters)) {
// 	            x := d.Letters[strings.ToUpper(string(c))]
// 	            out[i] += x.Lines[i]
// 	        } else if unicode.IsSpace(c) {
// 	            out[i] += strings.Repeat(" ", d.Length)
// 	        } else {
// 	            x := d.Letters["?"]
// 	            out[i] += x.Lines[i]
// 	        }
// 	    }
// 	}

// 	for i := 0; i < len(out); i++ {
// 		fmt.Println(out[i])
// 	}
// }

// func (d *Dictionnary) Print() {
// 	for d, e := range d.Letters {
// 		fmt.Printf("index: %s\n", d)
// 		e.Print()
// 	}
// }

// type Letter struct {
// 	Lines map[int]string
// }

// func (l *Letter) Print() {
// 	for i := 0; i < len(l.Lines); i++ {
// 		fmt.Println(l.Lines[i])
// 	}
// }

// func splitStr(s string, l int, b int, e int) (r []string) {
// 	i := b
// 	e = len(s)
// 	for i < e {
// 		j := i + l
// 		if j > e {
// 			j = e
// 		}
// 		r = append(r, string(s[i:j]))
// 		i = j
// 	}
// 	return r
// }

// func findStr(a string, list []string) bool {
//     for _, b := range list {
//         if b == a {
//             return true
//         }
//     }
//     return false
// }

// func Keys(m map[string]Letter) (keys []string) {
//     for k := range m {
//         keys = append(keys, k)
//     }
//     return keys
// }
