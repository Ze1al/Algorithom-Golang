/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/19
	字符串转换整数

	1. 如果第一个非空字符为正负号，则将该符号与之后尽可能连续起来
	2. 过滤整数后面的字符串
 */

package Top

import (
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	s = strings.TrimSpace(s)
	if string(s[0]) != "-" && !isNum(string(s[0])) {
		return 0
	}
	flag := false
	if string(s[0]) == "-"{
		s = s[1:]
		flag = true
	}
	if string(s[0]) == "+" {
		s = s[1:]
	}
	var stack []string
	for _, v := range s {
		if isNum(string(v)) {
			stack = append(stack, string(v))
		}else {
			break
		}
	}

	res := 0
	for i:=0; i<len(stack); i++ {
		num, _ := strconv.Atoi(stack[i])
		res = res * 10 + num
	}
	if flag {
		if res >= (1<<31) {
			return -1 * (1<<31)
		} else {
			return -1 * res
		}
	} else {
		if res >= (1<<31) {
			return (1<<31)
		} else {
			return res
		}
	}
}


func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

