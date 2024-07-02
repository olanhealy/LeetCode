# Notes

So just converting if from java over, was difficult to find a way to sort a string. Used chat to see about how can use rune slices, which are actually good for unicode characters too.  https://www.dotnetperls.com/convert-string-rune-slice-go


https://www.youtube.com/watch?v=XKyFc2fZuIg

# Complexity
> O(nlog (n)) using the sortString, the comparision of strings is O(n)

# Code
```Go
import ( 
    "fmt"
    "sort"
) 

func isAnagram(s string, t string) bool {

    sNew := sortString(s)
    tNew := sortString(t)

    if sNew == tNew {
        return true
    }
    return false
}

func sortString(s string) string {
	
    runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
```
