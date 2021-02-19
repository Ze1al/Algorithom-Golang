/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	最长的公共前缀
*/

package tencent

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		temp := ""
		for j := 0; j < len(ans) && j < len(strs[i]); j++ {
			if string(ans[j]) == string(strs[i][j]) {
				temp += string(ans[j])
			} else {
				break
			}
		}
		ans = temp
	}
	return ans
}
