package main

/*
 * HTTP命令行工具
 * Azuerc
 * github wsadiop123@qq.com
 */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Headers struct {
	RequestURL     string
	RequestMethod  string
	StatusCode     string
	RemoteAddress  string
	ReferrerPolicy string
}

type Querystring struct {
	Cc  string
	Ck  string
	Cl  string
	Ds  string
	Vl  string
	Et  string
	Ja  string
	Ln  string
	Lo  string
	Lt  string
	Rnd string
	Si  string
	V   string
	Lv  string
	Ct  string
	Tt  string
	Sn  string
}

type Jsonfile struct {
	Domainname  string
	Headers     Headers
	Querystring Querystring
	Post        string
}

func main() {
	var input string
	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)
	var err error
	var flag bool = true
	var jsfl Jsonfile
	for flag {
		fmt.Printf("http:")
		input, err = inputReader.ReadString('\n')
		if len(input) == 2 {
			input = "help"
		}
		var fs []string = strings.Fields(input)
		if err == nil {
			switch fs[0] {
			case "-a":
				httpall(fs, "GET")
				break
			case "-j":
				//-u C:\Users\Azuer\Desktop\golang\Task\test.json
				jsfl = getJsonfile(fs)
				break
			case "-u":
				updateJsonfile(fs, jsfl)
				break
			case "-q":
				flag = false
				break
			case "-help":
				fmt.Println("http:-a+请求网址 　#显示请求信息")
				fmt.Println("http:-j+json文件地址 　#显示文件信息")
				fmt.Println("http:-q           #退出")
				break
			default:
				fmt.Println("未知的命令,请输入-help查看使用方法!")
			}
		}
	}
}
func httpall(fs []string, method string) {

	req, err := http.NewRequest(method, "https://"+fs[1], nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Close = true
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println(string(body))
	fmt.Println("-------------------------------------------------------------------")
}

func getJsonfile(fs []string) Jsonfile {

	filep, err := os.Open(fs[1])
	if err != nil {
		fmt.Println(err)
	}
	defer filep.Close()

	var jsfile Jsonfile
	decoder := json.NewDecoder(filep)
	err = decoder.Decode(&jsfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("Domainname:" + jsfile.Domainname)
	fmt.Println("Headers:")
	fmt.Println("ReferrerPolicy:" + jsfile.Headers.ReferrerPolicy)
	fmt.Println("RemoteAddress:" + jsfile.Headers.RemoteAddress)
	fmt.Println("RequestMethod:" + jsfile.Headers.RequestMethod)
	fmt.Println("RequestURL:" + jsfile.Headers.RequestURL)
	fmt.Println("StatusCode:" + jsfile.Headers.StatusCode)
	fmt.Println("Querystring:")
	fmt.Println(jsfile.Querystring.Cc + "  " + jsfile.Querystring.Ck + "  " + jsfile.Querystring.Cl + "  " + jsfile.Querystring.Ds)
	fmt.Println(jsfile.Querystring.Vl + "  " + jsfile.Querystring.Et + "  " + jsfile.Querystring.Ja + "  " + jsfile.Querystring.Ln)
	fmt.Println(jsfile.Querystring.Lo + "  " + jsfile.Querystring.Lt + "  " + jsfile.Querystring.Rnd + "  " + jsfile.Querystring.Si)
	fmt.Println(jsfile.Querystring.V + "  " + jsfile.Querystring.Lv + "  " + jsfile.Querystring.Ct + "  " + jsfile.Querystring.Tt)
	fmt.Println(jsfile.Querystring.Sn)
	fmt.Println("Post:" + jsfile.Post)
	fmt.Println("-------------------------------------------------------------------")
	return jsfile
}
func updateJsonfile(fs []string, jsfl1 Jsonfile) {

}
