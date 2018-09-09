package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	fmt.Println("generete templates")

	if len(os.Args) < 2 {
		fmt.Println("input contest name")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("input contest number")
		return
	}

	// save template file path
	dp, _ := os.Getwd()
	tp := dp + "/templates.go"

	tmpl, err := os.Open("templates.go")
	if err != nil {
		fmt.Println("Cannot open template file")
		return
	}
	defer tmpl.Close()

	dirName := os.Args[1]
	conNum := os.Args[2]

	// move contest directory
	if _, e := os.Stat(dirName); os.IsNotExist(e) {
		if me := os.Mkdir(dirName, 0755); me != nil {
			fmt.Println("Cannot make contest directory")
			return
		}
	}
	if e := os.Chdir(dirName); e != nil {
		fmt.Println("Cannot move contest directory")
		return
	}

	// move target contest directory
	if _, e := os.Stat(conNum); os.IsNotExist(e) {
		if me := os.Mkdir(conNum, 0755); me != nil {
			fmt.Println("Cannot make number directory")
			return
		}
	}
	if e := os.Chdir(conNum); e != nil {
		fmt.Println("Cannot move number directory")
		return
	}

	pList := "abcdefghijklmnopqrstuvwxyz"

	// read generate file count
	fc := 4
	if len(os.Args) > 3 {
		tn, terr := strconv.Atoi(os.Args[3])
		if terr != nil {
			fmt.Println("input generate file count to args[3] ")
			return
		}
		fc = tn
	}

	// generate each file
	for i := 0; i < fc; i++ {
		gn := string(pList[i]) + ".go"
		if _, oe := os.Stat(gn); os.IsExist(oe) {
			continue
		}
		//read template file
		tmpl, err := os.Open(tp)
		if err != nil {
			fmt.Println("Cannot open template file")
			return
		}
		defer tmpl.Close()

		dst, cerr := os.Create(gn)
		if cerr != nil {
			fmt.Println("Cannot create", gn, "file")
		}
		defer dst.Close()

		if _, coerr := io.Copy(dst, tmpl); coerr != nil {
			fmt.Println("Cannot copy file")
		}

		fmt.Println("copy success", gn)
	}
}
