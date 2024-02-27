package threeSum_15

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

// hot 100 No.15 三数之和
/* [方法二]排序+双指针，双指针采用两头向中间逼近，保证[a,b,c]的a<=b<=c以及跳过相邻两次循环的相同元素来避免重复
 */
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums) // 先对数组从小到大排序，复杂度O(NlogN)
	l := len(nums)

	// 外层循环枚举a
	for i := 0; i < l-2; i++ {
		// 跳过重复的a
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 优化：如果最小的a都已经大于0了，那么后面的a+b+c一定大于0，此时直接跳出外层循环
		if nums[i] > 0 {
			break
		}

		// 优化：如果a加上最大的两个数都小于0了，那么该a的情况下三数和都会小于0，则continue继续搜索下一个更大的a
		if nums[i]+nums[l-1]+nums[l-2] < 0 {
			continue
		}

		target := -1 * nums[i] // b + c = target即可
		// 在固定a且满足b+c=target的情况下用双指针找到这个a下所有的b,c组合
		j, k := i+1, l-1 // 指针j从左往右枚举b，指针k从右往左枚举c
		for j < k {      // 两个指针向中间逼近，但保证b<=c
			// 跳过重复的b
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			// 跳过重复的c
			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			// 找到满足条件的一个三元组
			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 右移j的同时左移k，找到下一组跳过重复的b和c
				j++
				k--
				continue
			}

			// 若 b + c > target则将左移k（减小c，同时跳过相等的c）
			if nums[j]+nums[k] > target {
				k--
				continue
			}

			// 若 b + c < target则将右移j（增大b，同时跳过相等的b）
			if nums[j]+nums[k] < target {
				j++
				continue
			}
		}
	}
	return res
}

//func threeSum(nums []int) [][]int {
//	var res [][]int
//	sort.Ints(nums) // 先对数组从小到大排序，复杂度O(NlogN)
//	l := len(nums)
//
//	// 外层循环枚举a
//	for i := 0; i < l-2; i++ {
//		// 跳过重复的a
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//
//		// 优化：如果最小的a都已经大于0了，那么后面的a+b+c一定大于0，此时直接跳出外层循环
//		if nums[i] > 0 {
//			break
//		}
//
//		// 优化：如果a加上最大的两个数都小于0了，那么该a的情况下三数和都会小于0，则continue继续搜索下一个更大的a
//		if nums[i]+nums[l-1]+nums[l-2] < 0 {
//			continue
//		}
//
//		target := -1 * nums[i] // b + c = target即可
//		// 在固定a且满足b+c=target的情况下用双指针找到这个a下所有的b,c组合
//		j, k := i+1, l-1 // 指针j从左往右枚举b，指针k从右往左枚举c
//		for j < k {      // 两个指针向中间逼近，但保证b<=c
//			// 找到满足条件的一个三元组
//			if nums[j]+nums[k] == target {
//				res = append(res, []int{nums[i], nums[j], nums[k]})
//				// 右移j的同时左移k，找到下一组跳过重复的b和c
//				j++
//				k--
//				for j < k && nums[j] == nums[j-1] {
//					j++
//				}
//				for j < k && nums[k] == nums[k+1] {
//					k--
//				}
//				continue
//			}
//
//			// 若 b + c > target则将左移k（减小c，同时跳过相等的c）
//			if nums[j]+nums[k] > target {
//				k--
//				for j < k && nums[k] == nums[k+1] {
//					k--
//				}
//				continue
//			}
//
//			// 若 b + c < target则将右移j（增大b，同时跳过相等的b）
//			if nums[j]+nums[k] < target {
//				j++
//				for j < k && nums[j] == nums[j-1] {
//					j++
//				}
//				continue
//			}
//		}
//	}
//	return res
//}

/* [方法一]符合条件的三元组有三种情况：三个相同的数（三个零）、2同+1不同、3不同，找出nums中无重复的所有数，用双指针对其进行遍历，
外层循环负责找出2同+1不同的三元组，这种情况只要2同的数不等就不会重复。内层循环负责找出3不同的三元组，解决判断重复导致提交超时的
方法是将三元组的三个数排序后拼成string存入一个map，这样保证了三个同样的数只生成唯一的string，利用hash来快速判断该string是否
存在。尽管通过了测试，但是时间和空间消耗均较大。
*/
//func threeSum(nums []int) [][]int {
//	var res [][]int
//	existMp := map[string]int{} // 记录res中已经存在的属于三种不同数字的三元组string（去重用），如三元组[1, 0, -1]在其中存储为{"-101":1}
//	cntMp := map[int]int{}      // 数字 -> 出现次数
//	var keys []int              // nums中不重复的数字组成的切片（即cntMp的keys）
//	for _, num := range nums {
//		if _, exist := cntMp[num]; !exist {
//			keys = append(keys, num)
//		}
//		cntMp[num]++
//	}
//
//	// 三元组中为三个相同的数字（[0,0,0]）
//	if cntMp[0] >= 3 {
//		res = append(res, []int{0, 0, 0})
//	}
//
//	// 遍历keys
//	for i := 0; i < len(keys); i++ {
//		// 三元组中为2个相同+1个不同的数字的情况（[a,a,b]）
//		if cntMp[keys[i]] >= 2 && cntMp[-(keys[i]+keys[i])] >= 1 && keys[i] != 0 {
//			res = append(res, []int{keys[i], keys[i], -(keys[i] + keys[i])})
//		}
//
//		// 通过双指针找到符合条件的3个不同数字组成的三元组（[a,b,c]）
//		for j := i + 1; j < len(keys); j++ {
//			numK := -(keys[i] + keys[j]) // 第三个数
//			if cntMp[numK] > 0 && numK != keys[i] && numK != keys[j] {
//				threeNum := []int{keys[i], keys[j], numK}
//				sort.Ints(threeNum) // 对三个数进行排序
//				// 将排序完的三个数拼接成一个string，这样就能保证同样的三个数拼成的string唯一
//				str := strconv.Itoa(threeNum[0]) + strconv.Itoa(threeNum[1]) + strconv.Itoa(threeNum[2])
//
//				// 判断重复，用hash记录string可以保证查询已存在的三元组更快
//				if _, exist := existMp[str]; !exist {
//					res = append(res, []int{keys[i], keys[j], numK})
//					existMp[str] = 1
//				}
//			}
//		}
//	}
//	return res
//}

// 超时（怀疑是重复的三元组太多）
//func threeSum(nums []int) [][]int {
//	var res [][]int
//	existCntMps := []map[int]int{} // 记录res中已经存在三元组的次数映射（去重用）
//	cntMp := map[int]int{}         // 数字 -> 出现次数
//	for _, num := range nums {
//		cntMp[num]++
//	}
//
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			ok := true                   // 这三个数满足条件的标志
//			numK := -(nums[i] + nums[j]) // 第三个数
//			threeMp := map[int]int{}     // 这三个数中 数字 -> 出现次数
//			threeMp[nums[i]]++
//			threeMp[nums[j]]++
//			threeMp[numK]++
//
//			// 三个数均满足nums中个数的要求才行
//			for num, cnt := range threeMp {
//				if cnt > cntMp[num] {
//					ok = false
//					break
//				}
//			}
//
//			if !ok {
//				continue
//			}
//
//			// 若数量满足要求继续检查res中是否已有相同的三元组
//			dup := false // res中已经存在相同的三元组的标志
//			for _, mp := range existCntMps {
//				thisSame := true
//				for num, cnt := range threeMp {
//					if mp[num] != cnt { // 只要有某个num的数量不同就认为是不同的三元组
//						thisSame = false
//						break
//					}
//				}
//				if thisSame {
//					dup = true
//					break
//				}
//			}
//
//			// 如果不重复才加入res
//			if !dup {
//				res = append(res, []int{nums[i], nums[j], numK})
//				existCntMps = append(existCntMps, threeMp)
//			}
//
//		}
//	}
//	return res
//}
