package main

import (
	"fmt"
	"github.com/liboying/go-utils/cet"
)
/**:
 * @author: lby
 * @Desc:
 * @create: 2023-01-04 13:42
 */
func main() {
	fmt.Println("这是一个测试用的包")
	Method1()
}

func Method1() {
	fmt.Println("这是一个测试方法1")
	cet.Hello()
}
