package main

import (
	"fmt"
	"time"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/14     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/14 1:19 PM
 * @date 2019/10/14 1:19 PM
 * @since 1.0.0
 */
func main() {
	//var age int = 10
	//var a error
	tm := time.Now().UnixNano()
	//fmt.Println("test")
	fmt.Println(tm)
	fmt.Println(len(string(tm)))

	//fmt.Println(a.Error())
}
