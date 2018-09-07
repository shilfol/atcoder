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

////////////////////////////////////////
///        end templates             ///
////////////////////////////////////////
func intAbs(n int) int {

	return int(math.Abs(float64(n)))
}

func main() {
	line := nextLine()

	spl := intSprit(line)

	x0 := spl[0]
	y0 := spl[1]
	x1 := spl[2]
	y1 := spl[3]

	xd := intAbs(x0 - x1)
	yd := intAbs(y0 - y1)

	if x0 > x1 {
		if y0 > y1 {
			fmt.Println(x1+yd, y1-xd, x0+yd, y0-xd)
		} else {
			fmt.Println(x1-yd, y1-xd, x0-yd, y0-xd)
		}
	} else {
		if y0 > y1 {
			fmt.Println(x1+yd, y1+xd, x0+yd, y0+xd)
		} else {
			fmt.Println(x1-yd, y1+xd, x0-yd, y0+xd)
		}
	}
}
