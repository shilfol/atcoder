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

type tip struct {
	count int
	n     int
	sel   int
	six   []int
	nine  []int
}

func main() {
	line := readBigLine()
	N := parseInt(line)

	sixs := []int{6}
	nines := []int{9}

	for i := 0; i < 10; i++ {
		tmp := sixs[i] * 6
		if tmp > N {
			break
		}
		sixs = append(sixs, tmp)
	}
	for i := 0; i < 10; i++ {
		tmp := nines[i] * 9
		if tmp > N {
			break
		}
		nines = append(nines, tmp)
	}

	tips := []tip{{0, N, 6, sixs, nines}, {0, N, 9, sixs, nines}}

	n := 100000
	count := 100000

	var t tip
	for len(tips) > 0 {
		t, tips = tips[len(tips)-1], tips[:len(tips)-1]

		if t.n < t.sel {
			if n+count > t.n+t.count {
				n = t.n
				count = t.count
			}
			continue
		}
		if count < t.count {
			continue
		}

		num := 0
		if t.sel == 6 {
			num = t.six[len(t.six)-1]
			if t.n-num >= 0 {
				t.n -= num
				t.count++
			}
			if t.n < num {
				t.six = t.six[:len(t.six)-1]
			}
		} else {
			num = t.nine[len(t.nine)-1]
			if t.n-num >= 0 {
				t.n -= num
				t.count++
			}
			if t.n < num {
				t.nine = t.nine[:len(t.nine)-1]
			}
		}

		if len(t.six) > 0 {
			tips = append(tips, tip{t.count, t.n, 6, t.six, t.nine})
		}
		if len(t.nine) > 0 {
			tips = append(tips, tip{t.count, t.n, 9, t.six, t.nine})

		}
	}

	fmt.Println(n + count)

}
