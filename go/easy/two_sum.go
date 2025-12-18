func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for i, n := range nums {
		c := target - n
		if index, f := nMap[c]; f {
			return []int{index, i}
		}
		nMap[n] = i
	}
	return []int{}
}
