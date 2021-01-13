/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package train

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recursion(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}


func recursion(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recursion(A.Left, B.Left) && recursion(A.Right, B.Right)
}
