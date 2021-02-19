/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	字符串相乘
 */

package tencent

import "strconv"

func multiply(num1 string, num2 string) string {
	number1, _ := strconv.ParseFloat(num1, 64)
	number2, _ := strconv.ParseFloat(num1, 64)
	res := number1 * number2
	return strconv.FormatInt(int64(res), 64)
}