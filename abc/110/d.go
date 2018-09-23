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
func divs(n int) int {
	if n%2 == 0 {
		return 2
	} else {
		for i := 3; i < n; i += 2 {
			if n%i == 0 {
				return i
			}
		}
	}
	return -1
}

func main() {
	line := nextLine()

	spl := strSprit(line)

	N := parseInt(spl[0])
	M := parseInt(spl[1])

	fmt.Println(divs(N))
	fmt.Println(divs(M))

}
