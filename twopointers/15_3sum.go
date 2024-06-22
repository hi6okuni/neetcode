package twopointers

import (
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Example 2:

// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// Example 3:

// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

// Constraints:

// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105

// [-4, -1, -1, 0, 1, 2]
// [-4 -3 -2 -1 -1 0 0 1 2 3 4]

// タイムアウトのsolution
// func ThreeSum(nums []int) [][]int {
// 	ans := [][]int{}
// 	sumPool := map[int]map[[2]int]struct{}{}

// 	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

// 	for l := 0; l < len(nums)-2; l++ {
// 		for r := l + 1; r < len(nums)-1; r++ {
// 			sum := nums[l] + nums[r]

// 			pair := [2]int{nums[l], nums[r]}
// 			if pairs, ok := sumPool[sum]; ok {
// 				if _, pairExists := pairs[pair]; pairExists {
// 					continue
// 				}
// 			} else {
// 				sumPool[sum] = make(map[[2]int]struct{})
// 			}
// 			sumPool[sum][pair] = struct{}{}

// 			if sum+nums[len(nums)-1] < 0 {
// 				continue
// 			} else {
// 				for m := len(nums) - 1; m > r; m-- {
// 					if sum+nums[m] == 0 {
// 						ans = append(ans, []int{nums[l], nums[r], nums[m]})
// 						break
// 					}
// 				}
// 			}

// 		}
// 	}

// 	return ans
// }

func ThreeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums) // 配列をソート

	n := len(nums)

	for i := 0; i < n-2; i++ {
		// 重複を避けるためのチェック
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 2つポインタを利用して三つ組を探す
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 重複を避けるため、left と rightを移動
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++ // 合計が低い場合、leftを右に移動
			} else {
				right-- // 合計が高い場合、rightを左に移動
			}
		}
	}

	return ans
}
