package main

import (
	"fmt"
	"os"

	t "github.com/shilfol/atcoderTools/tools"
)

func main() {
	// go run trypost.go contestdir number diff file

	if len(os.Args) < 5 {
		fmt.Println("format")
		fmt.Println("go run trypost.go contestdir number diff file")
		return
	}

	con := os.Args[1] + os.Args[2]
	diff := os.Args[3]
	file := os.Args[4]

	t.TrySolve(con, diff, file)
}
