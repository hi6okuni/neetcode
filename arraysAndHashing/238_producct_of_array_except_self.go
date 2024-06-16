package arraysandhashing

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

// fromLeft: [1,1,2,6]
// fromRight: [24,12,4,1]
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)

	fromLeft := make([]int, length)
	fromRight := make([]int, length)

	for i := 0; i < length; i++ {
		if i == 0 {
			fromLeft[i] = 1
		} else {
			fromLeft[i] = fromLeft[i-1] * nums[i-1]
		}
	}

	fromRight[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		fromRight[i] = fromRight[i+1] * nums[i+1]
	}

	for i := 0; i < length; i++ {
		ans[i] = fromLeft[i] * fromRight[i]
	}

	return ans
}
