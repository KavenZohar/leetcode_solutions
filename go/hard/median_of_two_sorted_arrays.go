func findMedianSortedArrays(nums1 []int, nums2 []int) float64 { // [1, 3], [2]
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1 // nums1 = [2], nums2 = [1,3]
	}

	l := 0
	r := len(nums1)                          // 1
	mid := (len(nums1) + len(nums2) + 1) / 2 // (1+2+1) / 2 = 2
	for l <= r {
		i := (l + r) / 2 // loop 1: i = 0, loop 2: i = 1
		j := mid - i     // loop 1: j = 2, loop 2: j = 1

		l1 := -1 << 20
		if i > 0 {
			l1 = nums1[i-1] // loop 1: l1 = -100, loop 2: l1 = 2
		}

		r1 := 1 << 20
		if i < len(nums1) {
			r1 = nums1[i] // loop 1: r1 = 2, loop 2: r1 = 100
		}

		l2 := -1 << 20
		if j > 0 {
			l2 = nums2[j-1] // loop 1: l2 = 3, loop 2: l2 = 1
		}

		r2 := 1 << 20
		if j < len(nums2) {
			r2 = nums2[j] // loop 1: r2 = 100, loop 2: r2 = 3
		}

		if l1 <= r2 && l2 <= r1 { // loop 1: -100 <= 100 & 3 <= 2, loop 2: 2 <= 3 & 1 <= 100
			fmt.Println("ok")
			if (len(nums1)+len(nums2))%2 == 0 { // 3 % 2 = 1
				ml := l1
				if l2 > ml {
					ml = l2
				}

				minR := r1
				if r2 < minR {
					minR = r2
				}

				return float64(ml+minR) / 2.0

			}
			if l1 > l2 {
				return float64(l1) // 2.0
			}
			return float64(l2)
		}

		if l1 > r2 {
			r = i - 1
		} else {
			l = i + 1 // l = 1
		}
	}
	fmt.Println("not ok")
	return float64(-1)
}
