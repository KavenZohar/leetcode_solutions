func lengthOfLongestSubstring(s string) int {
	rs := 0
	for i, _ := range s {
		sj := s[i:]
		sMap := make(map[rune]int)
		a := 0
		for j, srn := range sj {
			if _, f := sMap[srn]; f {
				break
			}
			sMap[srn] = j
			a += 1
			if a > rs {
				rs = a
			}
		}
	}
	return rs
}
