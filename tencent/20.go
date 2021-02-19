/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	有效的括号
*/

package tencent

func isValid(s string) bool {
	stack :=[]string{}
	for i := 0; i < len(s); i++ {
		if isLeftParent(string(s[i])) {
			stack = append(stack, string(s[i]))
		} else {
			left := getParent(string(s[i]))
			if len(stack) > 0 && left == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func getParent(s string) string {
	switch s {
	case ")":
		return "("
	case "}":
		return "{"
	case "]":
		return "["
	}
	return ""
}

func isLeftParent(s string) bool {
	if s == "[" || s == "{" || s == "(" {
		return true
	}
	return false
}
