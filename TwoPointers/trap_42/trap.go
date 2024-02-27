package trap_42

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// hot 100 No.42 接雨水
// [方法三] 相向双指针，当前缀最大高度小于右边的后缀最大高度时，当前列能接水量取决于前缀最大高度和当前柱子高度，之后将左指针右移并更新前缀最大高度
// 当前缀最大高度大于右边的后缀最大高度时，当前列能接水量取决于后缀最大高度和当前柱子高度，之后将右指针左移并更新后缀最大高度
// 时间复杂度 O(n)，空间复杂度 O(1)
func trap(height []int) int {
	res := 0
	length := len(height)
	leftMax, rightMax := height[0], height[length-1] // 当前柱子左右两边最高的柱子高度(前缀、后缀最大高度，不包含当前柱子)
	l, r := 1, length-2                              // 左右指针

	// 计算每一列的接水量
	for l <= r {
		if leftMax < rightMax { // 当前缀最大高度小于右边的后缀最大高度时，当前列能接水量取决于前缀最大高度和当前柱子高度
			if height[l] < leftMax { // 该列能接到水
				res += leftMax - height[l]
			}
			leftMax = max(leftMax, height[l]) // 更新前缀最大高度
			l++                               // 左指针右移
		} else {
			if height[r] < rightMax {
				res += rightMax - height[r]
			}
			rightMax = max(rightMax, height[r]) // 更新后缀最大高度
			r--
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// [方法二] 按列算
//func trap(height []int) int {
//	res := 0
//	leftMax, rightMax := 0, 0 // 当前柱子左右两边最高的柱子高度
//
//	for cur := 1; cur < len(height)-1; cur++ {
//		if height[cur-1] > leftMax { // cur向右移则左边每次多出上一次的柱子
//			leftMax = height[cur-1]
//		}
//
//		if height[cur] >= rightMax { // 之前右边最高的可能就是cur，因此需要重新找到右边最高
//			max := 0
//			for i := cur + 1; i < len(height); i++ {
//				if height[i] > max {
//					max = height[i]
//				}
//			}
//			rightMax = max
//		}
//
//		// 只有当前列的柱子高度小于左右两边最高的柱子的较矮柱子时这一列才能接到水
//		if height[cur] < min(leftMax, rightMax) {
//			res += min(leftMax, rightMax) - height[cur]
//		}
//	}
//	return res
//}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//// [方法一] 双指针+分治法，首先将双指针分别定位到最外围的两个柱子（高度第一次开始下降则找到最外围的），然后不断调用subTrap()函数查看是否需要继续划分。
//// 划分点通过l和r之间的最高的柱子来确定，如果之间最高的柱子高度比两端的都小则证明不许要继续划分了，直接计算l和r作为两端柱子时的接水量。时间和空间复杂度
//// 均较高，需要另寻优解。
//func trap(height []int) int {
//	l, r := 0, len(height)-1
//
//	// 要接雨水最少要三根柱子（包括高度为0的）
//	if len(height) <= 2 {
//		return 0
//	}
//
//	// 两个指针向中间逼近，找到能接水的最外围的两个柱子（极大值点）
//	for l+1 <= r-1 {
//		move := false                 // 本轮是否有指针移动的标志
//		if height[l+1] >= height[l] { // 开始减小则找到最外围的左柱，否则l继续右移
//			l++
//			move = true
//		}
//
//		if height[r-1] >= height[r] { // 开始减小则找到最外围的右柱，否则r继续左移
//			r--
//			move = true
//		}
//
//		if !move { // 找到了最外围的左右柱子就退出循环
//			break
//		}
//	}
//
//	// 接不到雨水的情况
//	if l+1 > r-1 {
//		return 0
//	}
//
//	return subTrap(height, l, r)
//}
//
//// 计算左边或右边子区域的接水量
//// 不断划分子问题
//func subTrap(height []int, l int, r int) int {
//	// l和r已经挨着了，无法接水
//	if l >= r-1 {
//		return 0
//	}
//
//	lesser := height[l] // l和r中的较小高度
//	if height[r] < height[l] {
//		lesser = height[r]
//	}
//
//	mid := l + 1              // 用mid指针搜索l、r之间的划分点
//	betweenMax := height[mid] // l、r之间的最高的柱子高度
//	MaxIdx := mid             // l、r之间的最高的柱子的下标
//
//	// 搜素l、r之间找到最大高度的柱子
//	for mid <= r-1 {
//		if height[mid] > betweenMax {
//			betweenMax = height[mid]
//			MaxIdx = mid
//		}
//		mid++
//	}
//
//	// 需要划分计算两边的子区域
//	if betweenMax >= lesser {
//		// 分治求解左右两个区域的接水量
//		return subTrap(height, l, MaxIdx) + subTrap(height, MaxIdx, r)
//	}
//
//	// 若betweenMax < lesser，则l和r之间无划分点，直接计算l和r围城的接水量
//	res := lesser * (r - l - 1)
//	for i := l + 1; i < r; i++ {
//		res = res - height[i]
//	}
//	return res
//}
