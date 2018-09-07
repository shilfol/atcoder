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

func check(m [][]string, i, j int) int {
	n := make([]int, 4)

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for k := 0; k < 4; k++ {
		c := 0
		for {
			if m[i+dy[k]*(c+1)][j+dx[k]*(c+1)] != "." {
				n[k] = c
				break
			}
			c++
		}
	}

	return (n[0] + n[2]) * (n[1] + n[3])
}

func main() {
	line := nextLine()
	spl := strSprit(line)

	N := parseInt(spl[0])
	M := parseInt(spl[1])

	m := make([][]string, N+2)

	m[0] = make([]string, M+2)
	m[N+1] = make([]string, M+2)

	for i := 1; i <= N; i++ {
		m[i] = make([]string, M+2)
		l := nextLine()
		for j := 1; j <= M; j++ {
			m[i][j] = string(l[j-1])
		}
	}

	c := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if m[i][j] != "." {
				continue

			}

			c += check(m, i, j)
		}
	}

	fmt.Println(c)
}
