/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
	将字符串转成数字
 */

package tencent

import (
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	flag := false
	if len(s) == 0 {
		return 0
	}
	if string(s[0]) == "-" {
		flag = true
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	if !isNum(string(s[0])) {
		return 0
	}

	stack := []string{}
	for i:=0; i<len(s); i++ {
		if isNum(string(s[i])) {
			stack = append(stack, string(s[i]))
		} else {
			break
		}
	}

	res := 0
	for _, v := range stack {
		temp, _ := strconv.Atoi(v)
		res = res * 10 + temp
	}

	if flag {
		if res >= (1<<31) {
			return -1 * (1<<31)
		} else {
			return -1 * res
		}
	} else {
		if res >= (1<<31) {
			return 1<<31
		} else {
			return res
		}
	}


}


func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

