# Notes
easy enoguh to translate in go
Just the for loop starting from index 1 in the array looked at this https://yourbasic.org/golang/for-loop/

# Complexity
> O(n) as it has to go through each element

# Code
```Go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0;
    }
    var nonDuplicate int = 1
    

    for i := 1;  i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[nonDuplicate] = nums[i]
            nonDuplicate++
        }

        
    }

    return nonDuplicate
    
}
```
