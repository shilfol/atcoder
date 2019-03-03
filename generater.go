package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var tp string

func init() {
	// save template file path
	dp, _ := os.Getwd()
	tp = dp + "/template/templates.go"
}

func genereteFlie(fn string) error {
	//read template file
	tmpl, err := os.Open(tp)
	if err != nil {
		fmt.Println("Cannot open template file")
		return err
	}
	defer tmpl.Close()

	dst, cerr := os.Create(fn)
	if cerr != nil {
		fmt.Println("Cannot create", fn, "file")
		return cerr
	}
	defer dst.Close()

	if _, coerr := io.Copy(dst, tmpl); coerr != nil {
		fmt.Println("Cannot copy file")
		return coerr
	}
	return nil
}

func genereteFileSet(cnt int) error {
	pList := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < cnt; i++ {
		gn := string(pList[i]) + ".go"
		if _, oe := os.Stat(gn); os.IsExist(oe) {
			continue
		}

		if gerr := genereteFlie(gn); gerr != nil {
			return gerr
		}

		fmt.Println("copy success", gn)
	}
	return nil
}

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
	if gerr := genereteFileSet(fc); gerr != nil {
		fmt.Println(gerr)
		return
	}
}
