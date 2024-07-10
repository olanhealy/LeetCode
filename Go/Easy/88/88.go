func merge(nums1 []int, m int, nums2 []int, n int)  {
    var i = m -1
    var j = n -1
    var tracker = m + n -1

    for j >= 0 {
        if i >=0 && nums1[i] > nums2[j] {
            nums1[tracker] = nums1[i]
            i--
        } else {
            nums1[tracker] = nums2[j]
            j--
        }
        tracker--
    }

    
}
