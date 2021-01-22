/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/20
	罗马数字转成整数
 */

package Top

func romanToInt(s string) int {
	var res int
	pre := getInt(string(s[0]))
	for i := 1; i < len(s); i++ {
		temp := getInt(string(s[i]))
		if pre < temp {
			res -= pre
		} else {
			res += pre
		}
		pre = temp
	}
	res += pre
	return res
}

func getInt(ch string) int {
	switch ch {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	};
}