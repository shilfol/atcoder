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
type edge struct {
	start, end int
}

func dfs(n int, edges []edge, visited []bool) bool {
	stack := []int{}
	stack = append(stack, n)
	s := -1

	for len(stack) > 0 {
		s, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited[s] = true

		for _, e := range edges {
			if e.start == s {

				if visited[e.end] {
					return false
				}
				stack = append(stack, e.end)
			}
		}
	}
	return true
}

func main() {
	line := nextLine()
	spl := strSprit(line)
	N := parseInt(spl[0])
	M := parseInt(spl[1])

	visited := make([]bool, N+1)
	edges := make([]edge, M)

	for i := 0; i < M; i++ {
		line = nextLine()
		spl = strSprit(line)

		edges[i].start = parseInt(spl[0])
		edges[i].end = parseInt(spl[1])
	}

	count := 0

	for i := 1; i <= N; i++ {
		if visited[i] {
			continue
		}
		if dfs(i, edges, visited) {
			count++
		}
	}

	fmt.Println(count)

}
