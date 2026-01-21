// 3315. Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/description
package main

import "fmt"

func main() {

}

func minBitwiseArray(nums []int) []int {
	for i, p := range nums {
		removable := ((p + 1) & ^p) >> 1
		fmt.Printf("\tp: %08b (%d)", p, p)
		fmt.Printf("\t| p+1: %08b (%d)", p+1, p+1)
		fmt.Printf("\t| ^p: %08b (%d)", ^p, ^p)
		fmt.Printf("\t| (p+1)&^p: %08b (%d)", (p+1)&^p, (p+1)&^p)
		fmt.Printf("\t| removable: %08b (%d)\n", removable, removable)

		if removable == 0 {
			nums[i] = -1
			fmt.Println("\t| -1")
		} else {
			nums[i] = p ^ removable
			fmt.Printf("\t| %d\n", nums[i])
		}
	}
	return nums
}
