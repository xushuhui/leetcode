//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 1185 👎 0

package leetcode

import "testing"

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	j := 0
	for _, v := range nums {
		if v != 0 {
			nums[j] = v
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func moveZeroes1(nums []int) {
	l, left, right := len(nums), 0, 0
	for right < l {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++

	}
}
func Test_moveZeroes(t *testing.T) {
	nums := []int{
		0, 1, 0, 3, 12,
	}
	//moveZeroes(nums)
	moveZeroes1(nums)
}
