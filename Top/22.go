/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
	括号生成
	生成所有可能的并且有效的括号组合
	所有可能 一般都是回溯题
	递归生成所有可能情况
 */

package Top


func generateParenthesis(n int) []string {
	res := new([]string)
	context := ""
	brackets(context, res, n, n)
	return *res
}

func brackets(context string, res *[]string, left, right int) {
	if left == 0 && right == 0{
		*res = append(*res, context)
	}
	if left > right || left <0 || right < 0{
		return
	}
	brackets(context + "(", res, left - 1, right)
	brackets(context + ")", res, left, right - 1)
}

