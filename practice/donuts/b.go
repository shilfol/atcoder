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
	s := uint(i)

	return ((n >> s) & 1) == 1
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func matchCombo(i int, param []int) bool {
	c := 0

	for _, v := range param {
		if bitExist(i, v-1) {
			c++
		}
	}

	return c >= 3
}

func calcParam(n int, A []int, params [][]int) int {
	ret := 0
	for i := 0; i < len(A); i++ {
		if bitExist(n, i) {
			ret += A[i]
		}
	}

	for _, param := range params {
		if matchCombo(n, param[2:]) {
			ret += param[0]
		}
	}

	return ret
}

func main() {
	nums := nextLine()
	nspl := strSprit(nums)
	N := parseInt(nspl[0])
	M := parseInt(nspl[1])

	idol := nextLine()
	A := intSprit(idol)

	params := make([][]int, M)

	for i := 0; i < M; i++ {
		t := nextLine()
		params[i] = intSprit(t)
	}

	max := 0
	for i := 0; i < (1 << uint(N)); i++ {
		if bitCount(uint(i)) == 9 {
			ts := calcParam(i, A, params)
			if max < ts {
				max = ts
			}
		}
	}

	fmt.Println(max)
}
