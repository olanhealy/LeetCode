import "strings"

func strStr(haystack string, needle string) int {

if strings.Contains(haystack, needle) {
    return strings.Index(haystack, needle)
}
return -1

}
