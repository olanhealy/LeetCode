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
