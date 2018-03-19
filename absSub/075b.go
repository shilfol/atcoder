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

func strSprit(str, token string) []string {
	cols := strings.Split(str, token)
	return cols
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
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
	spl := strSprit(line, " ")
	H := parseInt(spl[0])
	W := parseInt(spl[1])

	maps := [][]string{}

	for i := 0; i < H; i++ {
		line = nextLine()
		mapline := strSprit(line, "")
		maps = append(maps, mapline)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if maps[i][j] != "." {
				continue
			}
			maps[i][j] = checkmine(maps, i, j, H, W)
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(maps[i][j])
		}
		fmt.Println()
	}
}

func checkmine(maps [][]string, i, j, H, W int) string {
	count := 0
	if i-1 >= 0 && maps[i-1][j] == "#" {
		count++
	}
	if i+1 < H && maps[i+1][j] == "#" {
		count++
	}
	if j-1 >= 0 && maps[i][j-1] == "#" {
		count++
	}
	if j+1 < W && maps[i][j+1] == "#" {
		count++
	}

	if i-1 >= 0 && j-1 >= 0 && maps[i-1][j-1] == "#" {
		count++
	}
	if i+1 < H && j-1 >= 0 && maps[i+1][j-1] == "#" {
		count++
	}
	if i-1 >= 0 && j+1 < W && maps[i-1][j+1] == "#" {
		count++
	}
	if i+1 < H && j+1 < W && maps[i+1][j+1] == "#" {
		count++
	}

	return strconv.Itoa(count)
}
