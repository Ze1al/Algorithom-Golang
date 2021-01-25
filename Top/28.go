/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
	实现strStr()
 */

package Top


func strStr(haystack string, needle string) int {
	L := len(haystack)
	n := len(needle)

	for i:=0; i<L-n+1;i++{
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}