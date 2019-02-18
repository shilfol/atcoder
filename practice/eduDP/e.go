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

func parseInt(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}

func intSprit(str string) []int64 {
	strs := strSprit(str)
	cols := make([]int64, len(strs))
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
func chmin(dp [][]int64, idx, dir int, dn int64) {
	if dp[idx][dir] > dn {
		dp[idx][dir] = dn
	}
}

func main() {
	line := nextLine()
	spl := intSprit(line)
	N := spl[0]
	W := spl[1]

	INF := int64(1 << 60)

	maxv := 100100
	dp := make([][]int64, N+10)
	for i := 0; i < int(N)+10; i++ {
		dp[i] = make([]int64, maxv)
		for j := 0; j < maxv; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0

	for i := 0; i < int(N); i++ {
		s := intSprit(nextLine())
		w, v := s[0], s[1]
		hv := int(v)
		for j := 0; j < maxv; j++ {
			if j-hv >= 0 {
				chmin(dp, i+1, j, dp[i][j-hv]+int64(w))
			}
			chmin(dp, i+1, j, dp[i][j])
		}
	}

	ret := 0
	for i := 0; i < maxv; i++ {
		if dp[N][i] <= W {
			ret = i
		}
	}
	fmt.Println(ret)

}
