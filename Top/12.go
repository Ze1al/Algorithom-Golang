/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/21
	贪心，每次使用最大值去匹配
*/

package Top

func intToRoman(num int) string {
	res := ""
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i< len(nums) && num >= 0; i++ {
		for nums[i] < num {
			num -= nums[i]
			res += getNumByRoman(nums[i])
		}
	}
	return res
}

func getNumByRoman(num int) string {
	switch num {
	case 1000:
		return "M"
	case 900:
		return "CM"
	case 500:
		return "D"
	case 400:
		return "CD"
	case 100:
		return "C"
	case 90:
		return "XC"
	case 50:
		return "L"
	case 40:
		return "XL"
	case 10:
		return "X"
	case 9:
		return "IX"
	case 5:
		return "V"
	case 4:
		return "IV"
	case 1:
		return "I"
	default:
		return ""
	}
}
