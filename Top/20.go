/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
	有效的括号
*/

package Top

var parent map[string]string = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
}

func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if isLeftParent(string(s[i])) {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) != 0 && parent[string(s[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0{
		return true
	} else {
		return false
	}
}


func isLeftParent(s string) bool {
	if s == "{" || s == "[" || s == "(" {
		return true
	}
	return false
}

