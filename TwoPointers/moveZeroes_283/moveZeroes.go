package moveZeroes_283

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// hot 100 No.283 移动零
func moveZeroes(nums []int) {
	var i, j int // 双指针，i指向已经处理到的下标，j负责找到第一个非零数
	found := false
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 { // 如果i指向的为0，则用j找到i以后的第一个不为0的数与i交换
			for j = i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					found = true
					break // 找到第一个非零数即推出内循环
				}
			}

			// 如果某一轮i后面没有找到非零数，则代表后面已经全零了，退出外循环
			if !found {
				break
			}

			found = false // 将found重置为false供下一轮用
		}
	}
}
