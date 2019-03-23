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

func setBit(d, n int) int {
	t := 1 << uint(n)
	return d | t
}

func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func try(g [][]bool, N int) bool {
	s := 0
	for i := 2; i <= N; i++ {
		if g[1][i] {
			s += i
		}
	}

	for i := 2; i <= N; i++ {
		ts := 0
		for j := 1; j <= N; j++ {
			if i != j && g[i][j] {
				ts += j
			}
		}
		if ts != s {
			return false
		}
	}
	return true
}

func printg(g [][]bool, N int) {
	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			if g[i][j] {
				fmt.Println(i, j)
			}
		}
	}
}

func main() {
	line := nextLine()
	N := parseInt(line)

	g := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		g[i] = make([]bool, N+1)
	}

	count := 0
	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			fmt.Println(i, j)
			count++
			g[i][j] = true
			g[j][i] = true

			if try(g, N) {
				fmt.Println(count)
				printg(g, N)
				return
			}
		}
	}
	fmt.Println(count)
}
