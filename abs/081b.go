package main

import (
	"bufio"
	"fmt"
	"math"
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
	N := parseInt(line)

	line = nextLine()
	spl := strSprit(line)

	count := math.MaxInt32

	for i := 0; i < N; i++ {
		num := parseInt(spl[i])
		tmpcount := 0
		for {
			if num%2 == 0 {
				num /= 2
				tmpcount++
			} else {
				break
			}
		}
		if tmpcount < count {
			count = tmpcount
		}
	}
	fmt.Println(count)
}
