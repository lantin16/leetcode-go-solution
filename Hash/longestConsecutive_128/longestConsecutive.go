package longestConsecutive_128

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// hot 100 No.128 最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mp := map[int]int{} // 存储整数到出现次数的映射
	for _, num := range nums {
		mp[num]++
	}

	cl := 1 // 当前找到的最长连续序列长度

	// 遍历map寻找最长连续序列的长度
	for k, v := range mp {
		if v == 0 {
			continue
		}

		// 若该数字的出现次数不为0
		tmpL := 1                         // 当前k所能找的连续序列长度
		mp[k] = 0                         // 遍历过的将出现次数归0
		for l := k - 1; mp[l] != 0; l-- { // 若向左能找到连续序列
			tmpL++
			mp[l] = 0
		}
		for r := k + 1; mp[r] != 0; r++ { // 若向右能找到连续序列
			tmpL++
			mp[r] = 0
		}

		if tmpL > cl {
			cl = tmpL
		}
	}
	return cl
}
