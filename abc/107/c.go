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

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////

func intAbs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

type pair struct {
	l, r int
	cp   int64
	t    int64
	c    int
}

func main() {
	line := readBigLine()
	spl := intSprit(line)

	N := spl[0]
	K := spl[1]

	line = readBigLine()
	nums := intSprit(line)

	sid := 0

	for i := 0; i < int(N); i++ {
		if nums[i] >= 0 {
			sid = i
			break
		}
	}

	plist := []pair{}

	if nums[sid] == 0 {
		plist = append(plist, pair{l: sid - 1, r: sid + 1, cp: 0, t: 0, c: 1})
	}

	plist = append(plist, pair{l: sid - 1, r: sid, cp: 0, t: 0, c: 0})

	mint := int64(100000000)
	p := pair{}
	for len(plist) > 0 {
		p, plist = plist[len(plist)-1], plist[:len(plist)-1]
		if p.c >= int(K) {
			if mint > p.t {
				mint = p.t
			}
			continue
		}
		nc := p.c + 1
		if mint < p.t {
			continue
		}

		if p.r < int(N) {
			plist = append(plist, pair{l: p.l, r: p.r + 1, cp: nums[p.r], t: p.t + intAbs(nums[p.r]-p.cp), c: nc})
		}
		if p.l >= 0 {
			plist = append(plist, pair{l: p.l - 1, r: p.r, cp: nums[p.l], t: p.t + intAbs(nums[p.l]-p.cp), c: nc})
		}
	}

	fmt.Println(mint)
}
