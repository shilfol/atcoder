package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func bitCount(n uint) int {

	x := uint64(n)

	const m = 1<<64 - 1
	const m0 = 0x5555555555555555
	const m1 = 0x3333333333333333
	const m2 = 0x0f0f0f0f0f0f0f0f

	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32

	return int(x) & (1<<7 - 1)
}

func bitExist(n, i int) bool {
	return ((n >> uint(i)) & 1) == 1
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func parseInt64(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}

func min64(nums ...int64) int64 {
	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = int64(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func intAbs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

func main() {
	line := nextLine()
	spl := intSprit(line)

	A := spl[0]
	B := spl[1]
	Q := spl[2]

	s := make([]int64, A)
	t := make([]int64, B)

	for i := 0; i < A; i++ {
		s[i] = parseInt64(nextLine())
	}
	for i := 0; i < B; i++ {
		t[i] = parseInt64(nextLine())
	}

	for i := 0; i < Q; i++ {
		x := parseInt64(nextLine())
		si := sort.Search(len(s), func(idx int) bool { return s[idx] >= x })
		ti := sort.Search(len(t), func(idx int) bool { return t[idx] >= x })

		res := int64(1 << 60)
		for n := si - 1; n <= si; n++ {
			if n < 0 || n >= len(s) {
				continue
			}
			for m := ti - 1; m <= ti; m++ {
				if m < 0 || m >= len(t) {
					continue
				}
				forward := intAbs(s[n]-x) + intAbs(s[n]-t[m])
				back := intAbs(t[m]-x) + intAbs(s[n]-t[m])
				res = min64(res, forward, back)
			}
		}
		fmt.Println(res)
	}
}
