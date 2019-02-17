package main

import (
	"bufio"
	"fmt"
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

func main() {
	line := nextLine()

	spl := intSprit(line)

	N := spl[0]

	m := make(map[string]int)
	b := make(map[string]bool)

	m["1"] = 2
	m["2"] = 5
	m["3"] = 5
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 3
	m["8"] = 7
	m["9"] = 6

	a := strSprit(nextLine())
	for _, v := range a {
		b[v] = true
	}
	sort.Strings(a)

	r := ""
	n := N

	if b["1"] {
		c := n / 2
		for i := 0; i < c; i++ {
			r += "1"
		}
		if n%2 == 0 {
			fmt.Println(r)
			return
		}
		switch {
		case b["7"]:
			r = "7" + r[1:]
		case b["5"]:
			r = "5" + r[2:]
		case b["3"]:
			r = "3" + r[2:]
		case b["2"]:
			r = "2" + r[2:]
		case b["8"]:
			r = "8" + r[3:]
		}
		fmt.Println(r)
		return
	}

	if b["7"] {
		c := n / 3
		for i := 0; i < c; i++ {
			r += "7"
		}
		if n%3 == 0 {
			fmt.Println(r)
			return
		}

		switch {
		case b["5"]:
			r = "5" + r[2:]
		case b["3"]:
			r = "3" + r[2:]
		case b["2"]:
			r = "2" + r[2:]
		case b["8"]:
			r = "8" + r[2:]
		}
		fmt.Println(r)
		return
	}

	for n > 0 {
		//pick
		d := 10
		s := ""
		for _, v := range a {
			if b[v] && m[v] <= d && s < v {
				d = m[v]
				s = v
			}
		}

		counts := n / d
		for i := 0; i < counts; i++ {
			r += s
		}
		n -= counts * d

		fmt.Println(d, s, r)
		if n != 0 {
			for n < d {
				r = r[:len(r)-1]
				n += d
			}
		}
		b[s] = false
	}
	fmt.Println(r)
}
