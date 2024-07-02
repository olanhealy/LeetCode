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
