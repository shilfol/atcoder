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
	return ((n >> uint(i)) & 1) == 1
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func chmax(dp []int, idx, n int) {
	if dp[idx] < n {
		dp[idx] = n
	}
}

func main() {
	line := nextLine()
	spl := intSprit(line)

	N := spl[0]
	M := spl[1]

	deg := make([]int, N)
	G := make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = []int{}
	}

	for i := 0; i < M; i++ {
		s := intSprit(nextLine())
		x, y := s[0], s[1]
		x--
		y--
		G[x] = append(G[x], y)
		deg[y]++
	}

	que := []int{}
	for i, v := range deg {
		if v == 0 {
			que = append(que, i)
		}
	}

	dp := make([]int, N+100)
	var q int
	for len(que) > 0 {
		q, que = que[0], que[1:]
		for _, v := range G[q] {
			deg[v]--
			if deg[v] == 0 {
				que = append(que, v)
				chmax(dp, v, dp[q]+1)
			}
		}
	}

	res := 0
	for i := 0; i < N; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	fmt.Println(res)
}
