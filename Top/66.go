/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
 */

package Top


func plusOne(digits []int) []int {
	digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10
	c := (digits[len(digits)-1] + 1) / 10
	
	for i:=len(digits)-1; i>0; i-- {
		temp := digits[i] + c + 1
		digits[i] = temp % 10
		c = temp / 10
	}
	return digits
}