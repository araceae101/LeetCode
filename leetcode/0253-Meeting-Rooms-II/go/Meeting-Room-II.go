package leetcode

// Problem: https://leetcode.com/problems/meeting-rooms-ii/
// Author: Araceae , Date: 2021/05/19
// Solution: Heap
// Time Complexity: O(NlogN)
// Space Complexity: O(N)

import (
	"container/heap"
	"sort"
)

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Top() int            { return h[0] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail
}

func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 1 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	h := &Heap{}
	heap.Push(h, intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= h.Top() {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i][1])
	}
	return h.Len()
}
