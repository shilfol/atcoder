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

type SortSlice []int

func (s SortSlice) Len() int {
	return len(s)
}

func (s SortSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func reverseString(str string) string {
	buf := []rune(str)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////

func rev(d int) int {
	rs := strconv.Itoa(d)
	s := reverseString(rs)

	return parseInt(s)

}

type pair struct {
	i, j int
}

func setMap(l []pair, am map[pair]int, i int) {
	for _, p := range l {
		am[p] = i
	}
}

func cal(x, y int, am map[pair]int) int {
	if x == 0 || y == 0 {
		fmt.Println("zero")
		return -1
	}

	m := make(map[pair]bool)

	l := []pair{}

	for c := 0; ; c++ {
		if x == 0 || y == 0 {
			//fmt.Println("zero")
			setMap(l, am, -1)
			return -1
		}

		p := pair{x, y}
		l = append(l, p)

		if _, ok := m[p]; ok {
			//fmt.Println("roop", c)
			setMap(l, am, 1)
			return c
		} else {
			m[p] = true
		}
		if t, ok := am[p]; ok {
			setMap(l, am, t)
			return t
		}

		if x < y {
			x = rev(x)
		} else {
			y = rev(y)
		}

		if x < y {
			y = y - x
		} else {
			x = x - y
		}

		//fmt.Println(c, ":", x, y)
	}

	//fmt.Println("out")
	setMap(l, am, -1)
	return -1
}

func main() {
	line := nextLine()

	spl := strSprit(line)
	N := parseInt(spl[0])
	M := parseInt(spl[1])

	m := make(map[pair]int)

	c := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			x, y := i, j
			if x > y {
				x, y = y, x
			}
			p := pair{x, y}
			if _, ok := m[p]; ok {
				if m[p] > 0 {
					c++
				}
				//fmt.Println("short: ", x, y)
				continue
			}
			//fmt.Println("do: ", x, y)
			m[p] = cal(x, y, m)
			if m[p] > 0 {
				c++
			}
		}
	}

	fmt.Println(c)
}
