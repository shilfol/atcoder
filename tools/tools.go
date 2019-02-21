package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type UserInfo struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

var client *http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client = &http.Client{
		Jar: jar,
	}
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

func inputUserInfo() UserInfo {
	var info UserInfo
	infofile := ".userinfo.json"

	if _, fe := os.Stat(infofile); !os.IsNotExist(fe) {
		d, err := ioutil.ReadFile(infofile)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(d, &info); err != nil {
			log.Fatal(err)
		}
		fmt.Println("* read: .userinfo.json")
		return info
	} else {
		fmt.Println(fe)
	}

	fmt.Println("* input userinfo")
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Username: ")
	sc.Scan()
	info.Username = sc.Text()

	fmt.Print("Password: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	info.Password = string(pass)
	fmt.Println()

	mar, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(infofile, mar, 0600); err != nil {
		log.Fatal(err)
	}
	fmt.Println("* write: .userinfo.json")

	return info
}

func fetchSession(n string) {

	info := inputUserInfo()

	URL := "https://atcoder.jp/login"

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	ck := resp.Cookies()
	spl := strings.Split(ck[1].String(), " ")
	tidx := strings.Index(spl[0], "csrf_token")
	combtoken, _ := url.QueryUnescape(spl[0][tidx:])
	cidx := strings.Index(combtoken, ":")
	cfin := strings.Index(combtoken, "=")
	csrfToken := combtoken[cidx+1 : cfin+1]

	data := url.Values{}
	data.Add("csrf_token", csrfToken)
	data.Add("password", info.Password)
	data.Add("username", info.Username)

	pres, err := client.PostForm(URL, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pres)
	v, _ := ioutil.ReadAll(pres.Body)
	fmt.Println(string(v))

	defer resp.Body.Close()

	pdata := url.Values{}
	pdata.Add("data.TaskScreenName", "abc118_d")
	pdata.Add("data.LanguageId", "3013")
	pdata.Add("csrf_token", csrfToken)

	f, err := os.Open("./abc/118/d-dp.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fv, _ := ioutil.ReadAll(f)

	pdata.Add("sourceCode", string(fv))

	SURL := "https://atcoder.jp/contests/abc118/submit"

	sres, err := client.PostForm(SURL, pdata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sres)
	sv, _ := ioutil.ReadAll(sres.Body)
	fmt.Println(string(sv))
	defer sres.Body.Close()

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
