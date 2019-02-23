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
type Pair struct {
	p, y int
}

func main() {
	line := nextLine()

	spl := strSprit(line)
	N := parseInt(spl[0])
	M := parseInt(spl[1])

	indexes := make(map[Pair]int, M)
	years := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		years[i] = []int{}
	}

	for i := 1; i <= M; i++ {
		s := intSprit(nextLine())
		years[s[0]] = append(years[s[0]], s[1])
		p := Pair{s[0], s[1]}
		indexes[p] = i
	}

	for i := 0; i <= N; i++ {
		sort.Ints(years[i])
	}

	ans := make([]string, M+1)

	for i := 1; i <= N; i++ {
		y := years[i]
		for idx, val := range y {
			tp := Pair{i, val}
			ai := indexes[tp]
			ans[ai] = fmt.Sprintf("%06d%06d", i, idx+1)
		}
	}
	for i := 1; i <= M; i++ {
		fmt.Println(ans[i])
	}

}
