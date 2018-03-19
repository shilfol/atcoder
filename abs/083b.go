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
	spl := strSprit(line)

	N := parseInt(spl[0])
	a := parseInt(spl[1])
	b := parseInt(spl[2])

	sum := 0
	for i := 1; i <= N; i++ {
		tmpsum := 0
		num := i
		for {
			tmpsum += num % 10
			if num/10 <= 0 {
				break
			}
			num /= 10
		}

		if a <= tmpsum && tmpsum <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}
