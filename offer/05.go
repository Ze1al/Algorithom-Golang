/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
 */

package offer

import "strings"

func replaceSpace(s string) string {
	res := strings.Replace(s, " ", "%20", -1)
	return res
}