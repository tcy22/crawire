package main

import (
	"regexp"
	"fmt"
)

const text  = `
              Hello tcy,you are the best tcy@123.com
              212121@lala.com 
              222222@haha.com hahahha
              `

func main()  {
	re := regexp.MustCompile(`[a-zA-z0-9]+@[a-zA-z0-9]+\.[a-zA-z0-9]+`)//用``不用转义
	match := re.FindString(text)//只找到一个,tcy@123.com
	fmt.Println(match)

	matchAll := re.FindAllString(text,-1)
	fmt.Println(matchAll)
	//-1代表全部返回[tcy@123.com 212121@lala.com 222222@haha.com]

	reg := regexp.MustCompile(`([a-zA-z0-9])+@([a-zA-z0-9])+\.([a-zA-z0-9])+`)//()起来的部分可以单独输出
	matchSub := reg.FindAllStringSubmatch(text,-1)//
	for _,m := range matchSub {
		fmt.Println(m)
	}
	//  [tcy@123.com y 3 m]
    //  [212121@lala.com 1 a m]
    //  [222222@haha.com 2 a m]
}
