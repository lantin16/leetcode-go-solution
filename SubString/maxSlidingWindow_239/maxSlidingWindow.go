package maxSlidingWindow_239

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。

// hot 100 No.239 滑动窗口最大值
// [方法二] 单调队列
// 用切片实现队列，用切片的取子切片操作来实现队列的弹出队首及插入队尾等操作
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var q []int // 记录严格单调递减的队列，存储的是对应元素在nums中的索引。这些下标按照从小到大的顺序被存储，并且它们在数组nums中对应的值是严格单调递减的。

	for i, x := range nums { // nums[i] = x
		// 1. 入
		// 将队尾比将要入队的元素x小或相等的元素全部出队，因为它们不可能成为最大值
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i) // 将x的索引入队

		// 2. 出
		// 如果队首元素已经不在滑动窗口中，则将其出队
		// 由于对每个元素都进行了检查，因此当队首元素出队后，队列中剩余元素就都在滑动窗口内了，所以这里只需要用if
		if i-k+1 > q[0] {
			q = q[1:]
		}

		// 3. 记录答案
		// 由于队列中的代表nums的数字从左到右严格单调递减，因此队首元素就是当前窗口的最大值
		if i >= k-1 { // 从k-1才开始记录答案，第一个窗口[0,k-1]
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// [暴力解法] 测试用例通过了，但超出时间限制。需要重新确定maxWd时检索新的滑动窗口最大值耗时过大。
//func maxSlidingWindow(nums []int, k int) []int {
//	var res []int
//	maxWd := nums[0] // 滑动窗口中的最大值
//	for i := 1; i < k; i++ {
//		if nums[i] > maxWd {
//			maxWd = nums[i]
//		}
//	}
//	res = append(res, maxWd)
//
//	for j := 0; j < len(nums)-k; j++ {
//		leave := nums[j]   // 要离开滑动窗口的数
//		enter := nums[j+k] // 要进入滑动窗口的数
//
//		// 如果加入的数比之前的maxWd都大，则它就是新窗口的最大值，直接更新maxWd
//		if enter >= maxWd {
//			maxWd = enter
//			res = append(res, maxWd)
//			continue
//		}
//
//		// 如果加入的数小于之前的maxWd，则进一步判断leave
//		// 如果leave不是maxWd，则maxWd不变
//		if leave < maxWd {
//			res = append(res, maxWd)
//			continue
//		}
//
//		// 如果leave是maxWd，则需要重新找到新窗口的最大值
//		maxWd = nums[j+1]
//		for r := j + 2; r <= j+k; r++ {
//			if nums[r] > maxWd {
//				maxWd = nums[r]
//			}
//		}
//		res = append(res, maxWd)
//	}
//	return res
//}
