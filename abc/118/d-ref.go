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

	//m["1"] = 2
	m["2"] = 5
	m["3"] = 5
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	//m["7"] = 3
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
		if n%3 == 1 {
			switch {
			case b["4"]:
				r = r[:len(r)-1] + "4"
			case b["9"]:
				r = "9" + r[2:]
			case b["6"]:
				r = "6" + r[2:]
			}
		}

		if n%3 == 2 {
			switch {
			case b["5"]:
				r = r[:len(r)-1] + "5"
			case b["3"]:
				r = r[:len(r)-1] + "3"
			case b["2"]:
				r = r[:len(r)-1] + "2"
			case b["8"]:
				r = "8" + r[2:]
			}
		}
		fmt.Println(r)
		return
	}

	//dfs
	type pair struct {
		n int
		r string
	}
	que := []pair{}
	for k, v := range m {
		if b[k] {
			que = append(que, pair{n - v, k})
		} else {
			delete(m, k)
		}
	}

	var q pair
	mr := ""
	for len(que) > 0 {
		q, que = que[0], que[1:]

		if q.n == 0 {
			if mr < q.r {
				mr = q.r
			}
			continue
		}

		for k, v := range m {
			if q.n-v >= 0 {
				que = append(que, pair{q.n - v, k + q.r})
			}
		}
	}
	fmt.Println(mr)
}
