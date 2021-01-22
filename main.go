/*
*  @author: didichuxing.com
 */


package main

import "fmt"

func main() {
	str := "Tom"
	fmt.Println(len(str))
	fmt.Println(str)
	fmt.Println(string(str[0]))
	fmt.Println(len([]rune(str)))
}
