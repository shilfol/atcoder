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
	line := readBigLine()
	N := parseInt(line)

	line = readBigLine()
	spl := strSprit(line)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = parseInt(spl[i])
	}

	counts := 0
	for bit := 0; bit < N; bit++ {
		bitsum := 0
		numsum := 0
		count := 0
		for i := bit; i < N; i++ {
			if count == 0 {
				numsum = a[i]
				bitsum = a[i]
			} else {
				numsum += a[i]
				bitsum ^= a[i]
			}
			count++
			if numsum == bitsum {
				fmt.Println(bit, bit+i)
				counts++
			}
		}
	}

	fmt.Println(counts)
}
