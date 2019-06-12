package twoSum

func twoSum(nums []int, target int) []int {
	compliments := make(map[int]int)
	for i, n := range nums {
		compliments[n] = i
	}

	for i, n := range nums {
		compliment := target - n
		compIdx, ok := compliments[compliment]
		if ok && i != compIdx {
			return []int{i, compIdx}
		}
	}
	return nil
}
