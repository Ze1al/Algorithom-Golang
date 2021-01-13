/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/7
	找到公共祖先
 */

package LeetCode


func lowestCommonAncestor(root, q, p *TreeNode) *TreeNode {
	// 保证p比较小
	if p.Val > q.Val {
		p, q = q, p
	}
	for root != nil {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > p.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}