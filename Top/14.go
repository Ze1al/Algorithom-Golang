/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/21
	最长公共前缀
*/

package Top

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	temp := strs[0]
	for i:=1; i<len(strs); i++ {
		res2 := ""
		for j:=0;j < len(strs[i])&&j<len(temp);j++ {
			if temp != "" && strs[i][j] == temp[j] {
				res2 += string(temp[j])
			}else{
				break
			}
		}
		temp = res2
	}
	return temp
}
