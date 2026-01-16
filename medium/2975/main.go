package main

import "sort"

func main() {

}

// Solution from AkitKumar on leetcode
// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/solutions/7498260/find-biggest-square-from-fence-gaps-begi-8o9m/
func AkitKumar(m int, n int, hFences []int, vFences []int) int {
	const MOD = 1000000007

	prep := func(cuts []int, limit int) []int {
		sort.Ints(cuts)
		out := []int{1}
		out = append(out, cuts...)
		out = append(out, limit)
		return out
	}

	h := prep(hFences, m)
	v := prep(vFences, n)

	gapSet := make(map[int]bool)

	for i := 0; i < len(h); i++ {
		for j := i + 1; j < len(h); j++ {
			gapSet[h[j]-h[i]] = true
		}
	}

	best := 0
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			d := v[j] - v[i]
			if d > best && gapSet[d] {
				best = d
			}
		}
	}

	if best == 0 {
		return -1
	}
	return int((int64(best) * int64(best)) % MOD)
}

// Solution from PattoBear on leetcode
// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/solutions/7498442/golang-good-ass-hints-by-pattobears-mq85/
func PattoBear(m int, n int, hFences []int, vFences []int) int {
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	hLen := len(hFences)
	vLen := len(vFences)
	mod := int(1e9 + 7)
	diffSet := map[int]struct{}{}

	for i := range hLen - 1 {
		for j := i + 1; j < hLen; j++ {
			diffSet[abs(hFences[i]-hFences[j])] = struct{}{}
		}
	}

	out := -1

	for i := range vLen - 1 {
		for j := i + 1; j < vLen; j++ {
			side := abs(vFences[i] - vFences[j])
			if _, ok := diffSet[side]; ok {
				out = max(out, side*side)
			}
		}
	}

	return out % mod
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
