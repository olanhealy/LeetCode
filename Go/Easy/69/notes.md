# Notes 
Easy to transverse to go, no interget overflow in go too so literally copy nd paste nearly
# Complexity
> O(log n)

# Code
```Go
func mySqrt(x int) int {
    if (x == 0 || x ==1) {
        return x
    }
    
    var start int = 1
    var end int = x
    var middle int

    for start <= end {
        middle = start + (end - start) / 2;

        var square int = middle * middle;

        if square == x {
            return middle
        } else if square < x {
            start = middle + 1
        } else {
            end = middle -1
        }
    }

    return end
}
```
