package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func strSprit(str string) []string {
	cols := strings.Split(str, " ")
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

func main() {
	line := nextLine()

	spl := strSprit(line)

	N := parseInt(spl[0])
	Y := parseInt(spl[1])

	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			k := N - i - j
			if k < 0 {
				break
			}
			if i*10000+j*5000+k*1000 == Y {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
