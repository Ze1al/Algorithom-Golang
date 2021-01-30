/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
	颜色分类
	对一个数组进行原地排序
	使相同的值相邻
	荷兰国旗问题
 */

package Top

// 单指针
func sortColors1(nums []int)  {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p += 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[p] = nums[p], nums[i]
			p += 1
		}
	}
}

// 双指针
func sortColors(nums []int)  {
	p1 := 0
	p2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1{
			nums[i], nums[p2] = nums[p2], nums[i]
			p2 ++
		} else if nums[i] == 0 {
			nums[i], nums[p1] = nums[p1], nums[i]
			if p1 < p2 {
				nums[i], nums[p2] = nums[p2], nums[i]
			}

			p1++
			p2++
		}
	}
}

















