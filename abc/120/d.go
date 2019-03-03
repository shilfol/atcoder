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
type UnionFind struct {
	Parent []int
}

func (uf *UnionFind) init(n int) {
	uf.Parent = make([]int, n+1)
	for i := 0; i <= n; i++ {
		uf.Parent[i] = -1
	}
}

func (uf *UnionFind) root(n int) int {
	if uf.Parent[n] < 0 {
		return n
	}
	uf.Parent[n] = uf.root(uf.Parent[n])
	return uf.Parent[n]
}

func (uf *UnionFind) issame(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) merge(x, y int) bool {
	x = uf.root(x)
	y = uf.root(y)

	if x == y {
		return false
	}
	if uf.Parent[x] > uf.Parent[y] {
		x, y = y, x
	}
	uf.Parent[x] += uf.Parent[y]
	uf.Parent[y] = x
	return true
}

func (uf *UnionFind) size(n int) int {
	return -uf.Parent[uf.root(n)]
}

func main() {
	line := nextLine()
	spl := intSprit(line)

	N, M := spl[0], spl[1]

	qa := make([]int, M)
	qb := make([]int, M)

	for i := 0; i < M; i++ {
		s := intSprit(nextLine())
		qa[i] = s[0]
		qb[i] = s[1]
	}

	var uf UnionFind
	uf.init(N)

	res := make([]int64, M)
	cur := (int64(N) * int64(N-1)) / (int64(2))

	for i := M - 1; i >= 0; i-- {
		res[i] = cur

		a, b := qa[i], qb[i]
		if uf.issame(a, b) {
			continue
		}

		sa, sb := int64(uf.size(a)), int64(uf.size(b))

		cur -= sa * sb
		uf.merge(a, b)
	}

	for i := 0; i < M; i++ {
		fmt.Println(res[i])
	}
}
