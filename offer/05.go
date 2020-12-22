/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/17
	替换空格
	s = "We are happy."
 */

package offer

import "strings"

func replaceSpace(s string) string {
	strings.Replace(s, " ", "%20", -1)
	return s
}