package main

import (
	"log"
	"time"
)

func main() {
	var test [][]int = [][]int{{1, 10, 4, 5}, {1, 2, 3, 6, 7}}

	for _, testex := range test {
		log.Println(twoSum(testex, 9))
	}
}

func twoSum(nums []int, target int) []int {
	var start = time.Now()
	defer func() {
		log.Printf("%d ns\n", time.Since(start).Nanoseconds())
	}()

	numbers := make(map[int]int)
	idxs := make(map[int]int)
	for idx, elem := range nums {
		numbers[elem] += 1
		idxs[elem] = idx
	}
	log.Println(numbers)
	log.Println(idxs)

	for idx, elem := range nums {
		if search := target - elem; numbers[search] > 0 && idx != idxs[search] {
			return []int{idx, idxs[search]}
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func val(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lhead := &ListNode{}
	ltail := lhead
	for true {
		sum := val(l1) + val(l2) + ltail.Val
		ltail.Val = sum % 10
		ltail.Next = &ListNode{Val: sum / 10}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if !(l1 != nil || l2 != nil) {
			if ltail.Next.Val == 0 {
				ltail.Next = nil
			}
			break
		}
		ltail = ltail.Next
	}
	return lhead
}
