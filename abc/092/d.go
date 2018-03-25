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

type SortSlice []int

func (s SortSlice) Len() int {
	return len(s)
}

func (s SortSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func reverseString(str string) string {
	buf := []rune(str)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////

func main() {
	line := nextLine()
	spl := strSprit(line)

	A := parseInt(spl[0])
	B := parseInt(spl[1])

	maps := [100][100]bool{}

	w, h := 100, 100

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j == 0 || (i%2 == 1 && j != w-1) {
				maps[i][j] = true
			}
		}
	}

	a, b := 1, 1

	i := 2
	j := 2

	for a < A {
		maps[i][j] = true
		a++
		j += 2

		if j > 98 {
			i += 2
			j = 2
		}
	}

	i = 97
	j = 97

	for b < B {
		maps[i][j] = false
		b++
		j -= 2

		if j < 2 {
			i -= 2
			j = 97
		}
	}

	fmt.Println(h, w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if maps[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
