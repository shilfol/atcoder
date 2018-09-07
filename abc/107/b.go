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

func blSprit(str string) []string {
	cols := strings.Split(str, "")
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

func main() {
	line := nextLine()
	spl := strSprit(line)

	H := parseInt(spl[0])
	W := parseInt(spl[1])

	r := make([][]string, H)
	o := [][]string{}
	fls := make([]bool, W)
	for i := 0; i < H; i++ {
		r[i] = make([]string, W)
		l := nextLine()
		r[i] = blSprit(l)
		fl := false

		for j := 0; j < W; j++ {
			if r[i][j] == "#" {
				fl = true
				fls[j] = true
			}
		}

		if fl {
			o = append(o, r[i])
		}
	}

	for n := 0; n < len(o); n++ {
		for i, f := range fls {
			if f {
				fmt.Print(o[n][i])
			}
		}
		fmt.Println()
	}

}
