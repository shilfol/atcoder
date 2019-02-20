package tools

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type ContestInfo struct {
	ContestName string `json:"contest_name"`
	User        string `json:"user"`
}

type Session struct {
	Session string `json:session`
}

func tryLogin() {
	sfileName := ".session.json"

	if _, fe := os.Stat(sfileName); os.IsNotExist(fe) {
		fetchSession(sfileName)
	}

	_, err := os.Open(sfileName)
	if err != nil {
		fmt.Println(err)
	}

}

func fetchSession(n string) {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Username: ")
	sc.Scan()
	name := sc.Text()

	fmt.Print("Password: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
	}
	sp := string(pass)

	fmt.Println()
	fmt.Println(name, sp)
}

func fetchTestcase() {

}
func doTestcase() {

}

func postAnswer() {

}

func TrySolve() {
	tryLogin()
	//fetchTestcase()
	//doTestcase()

	//postAnswer()

}
