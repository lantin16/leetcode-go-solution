package maxArea_11

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。

// hot 100 No.11 盛最多水的容器
// 双指针向中间逼近，时间消耗最少，优解
func maxArea(height []int) int {
	i, j := 0, len(height)-1                       // 左右双指针，起始指向两边的端点
	maxV := (j - i) * minInt(height[i], height[j]) // 最大容积，初始为两端点围成的

	// 双指针向中间缩（下面说的“缩“指的是板向中间移动而不是高度的缩小）
	// 经过分析，两板中矮的板若缩了后更矮则容积肯定变小，高的板缩了后无论变高还是变矮容积都变小
	// 因此只有矮的板缩了后变成高的板才可能出现更大的容积
	for i < j {
		// 先找到较矮的板
		if height[i] < height[j] { // 左边更矮
			i++                          // 将矮板右移
			if height[i] > height[i-1] { // 当右移后更高才比较容积
				maxV = maxInt(maxV, minInt(height[i], height[j])*(j-i))
			}
		} else { // 右边更矮
			j--                          // 将矮板左移
			if height[j] > height[j+1] { // 当左移后更高才比较容积
				maxV = maxInt(maxV, minInt(height[i], height[j])*(j-i))
			}
		}

	}
	return maxV
}

// 求两个int的最小值
func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 求两个int的最大值
func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 跳过某些外循环，可以通过，但耗时仍较长
//func maxArea(height []int) int {
//	var i, j int // 左右双指针
//	maxV := -1   // 最大容积
//	h := -1      // 当前最大容积的左板高度
//
//	for i = 0; i < len(height)-1; i++ {
//		// 每一次外循环都是以i为固定左板，移动右板来比较
//
//		// 若左板右移后高度比h大才可能出现最大容积
//		// 若左板向右移了还变低了则此趟必不可能出现最大容积，直接跳过
//		if height[i] > h {
//			for j = i + 1; j < len(height); j++ {
//				tmpH := height[i]
//				if tmpH > height[j] {
//					tmpH = height[j]
//				}
//				if tmpH*(j-i) > maxV {
//					maxV = tmpH * (j - i)
//					h = height[i] // 保证h始终为迄今为止找到的最大容积的左板高度
//				}
//			}
//		}
//	}
//	return maxV
//}

//// 暴力法会超出时间限制
//func maxArea(height []int) int {
//	var i, j int // 左右双指针
//	res := 0     // 最大容积
//
//	for i = 0; i < len(height)-1; i++ {
//		for j = i + 1; j < len(height); j++ {
//			h := height[i] // 水能达到的最大高度（取决于短板）
//			if height[i] > height[j] {
//				h = height[j]
//			}
//			if h*(j-i) > res {
//				res = h * (j - i)
//			}
//		}
//	}
//	return res
//}
