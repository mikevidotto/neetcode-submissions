func hasDuplicate(nums []int) bool {
    freq_map := make(map[int]int)
    for _, num := range nums {
        freq_map[num]++
    }

    for _, num := range freq_map {
        if num > 1 {
            return true 
        }
    }

    return false
}
