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
	return s[i] > s[j]
}

func main() {
	line := nextLine()
	N := parseInt(line)

	line = nextLine()
	spl := strSprit(line)

	nums := make(SortSlice, N)
	for i := 0; i < N; i++ {
		nums[i] = parseInt(spl[i])
	}
	sort.Sort(nums)

	alice, bob := 0, 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			alice += nums[i]
		} else {
			bob += nums[i]
		}
	}

	fmt.Println(alice - bob)
}
