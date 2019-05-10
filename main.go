package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 99; i++ {
		fmt.Println("![](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/course1/%E5%B9%BB%E7%81%AF%E7%89%87" + strconv.Itoa(i) + ".JPG)\n")
	}
}
