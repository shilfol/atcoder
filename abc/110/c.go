package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////
///        templates                 ///
////////////////////////////////////////

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readBigLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func strSprit(str string) []string {
	cols := strings.Split(str, " ")
	return cols
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func intSprit(str string) []int {
	strs := strSprit(str)
	cols := make([]int, len(strs))
	for i, v := range strs {
		cols[i] = parseInt(v)
	}
	return cols
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////

func main() {
	S := readBigLine()
	T := readBigLine()

	sm := map[rune]int{}
	ss := make([][]int, 26)

	sk := 0
	for i, s := range S {
		if im, ok := sm[s]; !ok {
			sm[s] = sk
			ss[sk] = []int{i}
			sk++
		} else {
			ss[im] = append(ss[im], i)
		}
	}

	tm := map[rune]int{}
	tk := 0
	var si int
	for i, t := range T {
		if _, ok := tm[t]; !ok {
			tm[t] = tk
			tk++
		}
		tmi := tm[t]

		if len(ss[tmi]) <= 0 {
			fmt.Println("No")
			return
		}
		si, ss[tmi] = ss[tmi][0], ss[tmi][1:]

		if si != i {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
