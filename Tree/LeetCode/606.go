/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/8
	根据二叉树创建字符串
 */

package LeetCode

import "strconv"

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val) + ""
	}
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")(" + tree2str(t.Right) + ")"
}
