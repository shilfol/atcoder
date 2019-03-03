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

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}

func setBit(d, n int) int {
	t := 1 << uint(n)
	return d | t

}

func main() {
	line := nextLine()
	spl := intSprit(line)

	N := spl[0]

	o := spl[1:]
	target := [][]int{}

	target = append(target, []int{o[0], o[1], o[2]})
	target = append(target, []int{o[0], o[2], o[1]})
	target = append(target, []int{o[1], o[0], o[2]})
	target = append(target, []int{o[1], o[2], o[0]})
	target = append(target, []int{o[2], o[0], o[1]})
	target = append(target, []int{o[2], o[1], o[0]})

	m := make(map[int]bool, 3)
	for _, v := range o {
		m[v] = true
	}

	uc := 0
	l := []int{}

	for i := 0; i < N; i++ {
		t := parseInt(nextLine())
		if m[t] {
			m[t] = false
			uc = setBit(uc, i)
		}
		l = append(l, t)
	}

	res := 10000000
	for _, mt := range target {

		tres := 0
		tuc := uc
		for _, v := range mt {
			if !m[v] {
				continue
			}
			mindiff := 100000
			minset := 0
			for i := 1; i < (1 << uint(N)); i++ {
				cf := false
				for j := 0; j < 8; j++ {
					if bitExist(tuc, j) && bitExist(i, j) {
						cf = true
						continue
					}
				}
				if cf {
					continue
				}

				tdiff := 0
				tsum := 0

				tdiff += (bitCount(uint(i)) - 1) * 10

				for t := 0; t < N; t++ {
					if bitExist(i, t) {
						tsum += l[t]
					}
				}

				tdiff += intAbs(tsum - v)

				if mindiff > tdiff {
					mindiff = tdiff
					minset = i
				}
			}
			tres += mindiff

			tuc = tuc | minset
		}
		if res > tres {
			res = tres
		}

	}
	fmt.Println(res)

}
