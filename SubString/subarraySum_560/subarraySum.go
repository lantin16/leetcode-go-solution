package subarraySum_560

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列

// hot 100 No.560 和为 K 的子数组
// [方法二] 前缀和+hash
func subarraySum(nums []int, k int) int {
	res := 0
	preMp := map[int]int{0: 1} // 前缀和 -> 出现次数，初始化前缀和0的次数为1
	// 这里为什么要这样初始化：因为若pre[j] == k了（即从头开始的前缀和已经等于k）则preMp[0]可能会少记录一次这种情况
	pre := 0 // 当前的前缀和

	// pre[j] - pre[i] == k  =>  pre[i] == pre[j] - k
	for j := 0; j < len(nums); j++ {
		pre += nums[j]
		res += preMp[pre-k]
		preMp[pre]++
	}

	return res
}

//// [方法一] 枚举法
//func subarraySum(nums []int, k int) int {
//	res := 0
//
//	// 先固定子串的起始点i
//	for i := 0; i < len(nums); i++ {
//		sum := 0
//		for j := i; j < len(nums); j++ {
//			sum += nums[j]
//			if sum == k {
//				res++
//			}
//		}
//	}
//
//	return res
//}
