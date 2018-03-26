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

func strmapSprit(str string) []string {
	cols := strings.Split(str, "")
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

type point struct {
	x, y int
}

func isVisit(s string, b bool) bool {
	return (s != "#") && !b

}

func main() {
	line := nextLine()
	spl := strSprit(line)
	H := parseInt(spl[0])
	W := parseInt(spl[1])

	maps := make([][]string, H)
	visited := make([][]bool, H)

	start := point{-1, -1}

	for i := 0; i < H; i++ {
		maps[i] = make([]string, W)
		visited[i] = make([]bool, W)
		line = nextLine()
		spl = strmapSprit(line)
		for j := 0; j < W; j++ {
			maps[i][j] = spl[j]
			if spl[j] == "s" {
				start = point{j, i}
			}
		}
	}

	stack := []point{}
	stack = append(stack, start)
	p := point{}

	for len(stack) > 0 {
		p, stack = stack[len(stack)-1], stack[:len(stack)-1]
		//fmt.Println(p)

		if p.x < 0 || p.x >= W || p.y < 0 || p.y >= H {
			//fmt.Println("out of range")
			continue
		}
		if !isVisit(maps[p.y][p.x], visited[p.y][p.x]) {
			//fmt.Println("not available")
			continue
		}

		visited[p.y][p.x] = true

		if maps[p.y][p.x] == "g" {
			fmt.Println("Yes")
			return
		}

		stack = append(stack, point{p.x, p.y + 1})
		stack = append(stack, point{p.x + 1, p.y})
		stack = append(stack, point{p.x, p.y - 1})
		stack = append(stack, point{p.x - 1, p.y})
	}

	fmt.Println("No")

}
