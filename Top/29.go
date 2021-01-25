/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
	两数相除
 */

package Top

func divide(dividend int, divisor int) int {
	flag := false
	if (dividend < 0 && divisor < 0){
		dividend = abs(dividend)
		divisor = abs(divisor)
		flag = false
	} else if dividend < 0 {
		dividend = abs(dividend)
		flag = true
	} else if divisor < 0 {
		divisor = abs(divisor)
		flag = true
	}

	if divisor > dividend {
		return 0
	}
	res := 1
	temp := divisor
	for dividend > temp+temp {
		res += res
		temp += temp
	}
	res += divide(dividend-temp, divisor)

	if flag {
		if res >= 1<<31 {
			return -1 * 1<<31
		}
		return -1 * res
	} else {
		if res >= 1<<31 {
			return  1<<31
		}
		return res
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

