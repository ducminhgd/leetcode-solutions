// 3494. Find the Minimum Amount of Time to Brew Potions
// https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/description
package main

func main() {

}

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	for i := 1; i < n; i++ {
		skill[i] += skill[i-1]
	}
	time := skill[n-1] * mana[m-1]
	for i := 1; i < m; i++ {
		maxtime := skill[0] * mana[i-1]
		for j := 1; j < n; j++ {
			currtime := skill[j]*mana[i-1] - skill[j-1]*mana[i]
			maxtime = max(maxtime, currtime)
		}
		time += maxtime
	}
	return int64(time)
}
