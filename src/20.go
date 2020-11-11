/*
*  @author: didichuxing.com
 */

package src


func isValid(s string) bool {
	if len(s) % 2 == 1 {
		return false
	}
	// 维护一个栈
	stack := []byte{}

	dict := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	for i:=0; i<len(s);i++{
		if dict[s[i]] > 0{
			if len(stack) == 0 || stack[len(stack)-1] != dict[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)==0
}



func _isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
