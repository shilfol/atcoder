package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N := parseInt(line)

	nums := make(SortSlice, N)
	for i := 0; i < N; i++ {
		line = nextLine()
		nums[i] = parseInt(line)
	}
	sort.Sort(nums)

	mocho := []int{}

	prev := 0
	for _, n := range nums {
		if prev != n {
			mocho = append(mocho, n)
			prev = n
		}
	}

	fmt.Println(len(mocho))
}
