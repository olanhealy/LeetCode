func runningSum(nums []int) []int {
    solution := make([]int, len(nums))
    solution[0] = nums[0]
    for i := 1; i < len(nums); i++ {
	solution[i] = solution [i - 1] + nums[i]  
    }
    return solution

}
