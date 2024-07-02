# Notes
Using hashmap in go. https://www.youtube.com/watch?v=p4LS3UdgJA4

What i used then trnasfer over what did in Java already

# Complexity
O(N) as it just depends on the size of the array

# Code
```Go
func containsDuplicate(nums []int) bool {
    place := make(map[int] bool)

    for _, i :=  range nums {
        if place[i] {
            return true
        }
        place[i] = true
    }
    return false
    
}
```
