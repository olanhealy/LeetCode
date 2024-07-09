# Notes
easy enough to convert to go
# Complexity
> O(n) depending on what int n is
# Code
```Go
func climbStairs(n int) int {
    if n <= 1 {
        return 1
    } else if n == 2 {
        return 2
    }

    var prev1 = 1
    var prev2 = 2
    var placeholder = 0
    
    for i := 3;  i <= n; i++ {
        placeholder = prev1 + prev2
        prev1 = prev2
        prev2 = placeholder 


    }
    return placeholder
    
}
```
