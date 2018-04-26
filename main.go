package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"regexp"
)

func main(){
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//解决乱码问题，将GBK转成utf8
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	if resp.StatusCode != http.StatusOK {//得到头部，body之类的
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}
	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}

//获取到r的encoding
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}

func printCityList(contents []byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents,-1)     //[][][]byte

	for _,m := range matches {
		fmt.Printf("City: %s , URL: %s\n",m[2],m[1])
	}
}