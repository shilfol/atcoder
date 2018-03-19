package main

import (
	"bufio"
	"fmt"
	"math"
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
	N := parseInt(line)

	var t [100000]int
	var x [100000]int
	var y [100000]int

	for i := 0; i < N; i++ {
		line = nextLine()
		spl := strSprit(line)
		t[i] = parseInt(spl[0])
		x[i] = parseInt(spl[1])
		y[i] = parseInt(spl[2])
	}

	px, py, pt := 0, 0, 0
	for i := 0; i < N; i++ {
		tdiff := t[i] - pt
		ldiff := int(math.Abs(float64(x[i]-px)) + math.Abs(float64(y[i]-py)))

		if tdiff < ldiff {
			fmt.Println("No")
			return
		}
		if tdiff%2 != ldiff%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}
