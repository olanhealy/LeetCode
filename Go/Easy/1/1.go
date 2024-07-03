func twoSum(nums []int, target int) []int {
    solution := make(map[int] int)

    for i, num := range nums {
        var x = target - nums[i]
         if index, found := solution[x]; found {
            return []int{index, i}
        }
        solution[num] = i
    }


    return nil
}
