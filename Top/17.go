/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/21
	电话号码的字母组合

	回溯：列举所有可能
*/

package Top

var phoneNumber map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = []string{}
	backtrack("", digits)
	return res
}

// 递归带货
func backtrack(conbination, next string) {
	if len(next) == 0 {
		res = append(res, conbination)
	} else {
		for _, v := range phoneNumber[string(next[0])] {
			backtrack(conbination+string(v), next[1:])
		}
	}
}
