package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

//从网络上获取数据的模块
func Fetcher(url string) ([]byte,error){
	<- rateLimiter
	resp,err := http.Get(url)

	//fmt.Printf("Fetcher resp url:%#v\n",resp)
	//fmt.Printf("Fetcher resp url:%s,resp.StatusCode:%d\n",url,resp.StatusCode)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {//得到头部，body之类的
		return nil,fmt.Errorf("rong status code: %d",resp.StatusCode)
	}

	//解决乱码问题，将GBK转成utf8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//获取到r的encoding
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}
