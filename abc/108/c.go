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

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
type pair struct {
	a, b int
}

func main() {
	line := nextLine()

	spl := intSprit(line)
	N := spl[0]
	K := spl[1]

	ps := []pair{}

	for i := 1; i <= N; i++ {
		for j := K - i%K; j <= N; j += K {
			ps = append(ps, pair{i, j})
		}
	}

	c := 0
	for _, p := range ps {
		for i := K - p.a%K; i <= N; i += K {
			if (p.b+i)%K == 0 {
				c++
			}
		}
	}
	fmt.Println(c)
}
