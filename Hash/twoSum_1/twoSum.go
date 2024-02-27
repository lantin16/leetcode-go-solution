package twoSum_1

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。

// hot 100 No.1 两数之和
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int) // 数字->它出现在nums的下标
	res := []int{}

	for i, num := range nums {
		// 插入前先检查mp中是否已有key为target-num，有的话直接得到结果并且可以保证不和自己匹配（可以和相等的另一个匹配）
		if v, ok := mp[target-num]; ok {
			res = append(res, i, v)
			break
		}
		mp[num] = i
	}

	return res
}
