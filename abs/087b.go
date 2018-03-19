package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func strSprit(str string) []string {
	cols := strings.Split(str, " ")
	return cols
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func main() {
	line := nextLine()
	a := parseInt(line)
	line = nextLine()
	b := parseInt(line)
	line = nextLine()
	c := parseInt(line)
	line = nextLine()
	x := parseInt(line)

	count := 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					count++
				}

			}
		}
	}

	fmt.Println(count)
}
