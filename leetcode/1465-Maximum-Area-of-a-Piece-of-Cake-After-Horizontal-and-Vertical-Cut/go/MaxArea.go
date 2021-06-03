package leetcode

/*
* Problem: https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
* Author: Araceae
* Date: 2021/06/03

* Solution: Slice Sort and Traverse
* to find the maximum long and width of the area

* Time Complexity: O(NlogN + MLogM)
* Space Complexity: O(1)

* Runtime: 84 ms, faster than 17.35% of Go online submissions for Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts.
* Memory Usage: 10 MB, less than 5.10% of Go online submissions for Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts.
 */

import "sort"

const modulo = 1<<9 + 7

// Problem: https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
func MaxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, []int{0, h}...)
	verticalCuts = append(verticalCuts, []int{0, w}...)

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxHoriCut, maxVertiCut := 0, 0

	for i := 1; i < len(horizontalCuts); i++ {
		maxHoriCut = max(maxHoriCut, horizontalCuts[i]-horizontalCuts[i-1])
	}

	for i := 1; i < len(verticalCuts); i++ {
		maxVertiCut = max(maxVertiCut, verticalCuts[i]-verticalCuts[i-1])
	}

	return maxHoriCut * maxVertiCut % modulo
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
