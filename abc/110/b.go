package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	line := nextLine()

	spl := strSprit(line)

	X := parseInt(spl[2])
	Y := parseInt(spl[3])

	xl := nextLine()
	x := intSprit(xl)
	sort.Ints(x)

	yl := nextLine()
	y := intSprit(yl)
	sort.Ints(y)

	if x[len(x)-1] >= y[0] {
		fmt.Println("War")
		return
	}

	if X >= y[0] {
		fmt.Println("War")
		return
	}
	if Y < y[0] {
		fmt.Println("War")
		return
	}

	fmt.Println("No War")
}
