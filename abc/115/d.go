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

func main() {
	line := nextLine()

	spl := strSprit(line)

	N := int64(parseInt(spl[0]))
	X, _ := strconv.ParseInt(spl[1], 10, 64)

	nums := make([]int64, N+1)
	ps := make([]int64, N+1)

	nums[0] = 1
	ps[0] = 1
	for i := 1; int64(i) <= N; i++ {
		nums[i] = nums[i-1]*2 + 3
		ps[i] = ps[i-1]*2 + 1
		fmt.Println(nums[i], ps[i], nums[i]-ps[i])
	}

	id := 0
	for i := 0; int64(i) <= N; i++ {
		if nums[i] >= X {
			fmt.Println(i, X)
			id = i
		}
	}
	r := int64(0)
	ret := int64(0)
	for {
		t := r + nums[id-1]
		if t < X {
			r += nums[id-1]
		} else if t == X {
			break
		}
		id--
		if id == 0 {
			break
		}
		fmt.Println(id, r)
	}
	fmt.Println(r)
}
