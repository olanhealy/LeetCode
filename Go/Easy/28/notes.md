# Notes
Easy to transverse to go just had search these https://pkg.go.dev/strings#example-Index

https://pkg.go.dev/strings#example-Contains

# Complexity
> O(m x n) as contains and index operations

# Code
```Go
import "strings"

func strStr(haystack string, needle string) int {

if strings.Contains(haystack, needle) {
    return strings.Index(haystack, needle)
}
return -1
    
}
```
