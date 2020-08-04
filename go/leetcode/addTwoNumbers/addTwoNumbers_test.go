package addTwoNumbers

import "testing"

type addTwoNumbersTestCase struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

var addTwoNumbersTestCases []addTwoNumbersTestCase = []addTwoNumbersTestCase{
	{nil, nil, &ListNode{Val: 0}},
	{&ListNode{Val: 0}, &ListNode{Val: 0}, &ListNode{Val: 0}},
	{&ListNode{Val: 0}, &ListNode{Val: 1}, &ListNode{Val: 1}},
	{&ListNode{Val: 1}, &ListNode{Val: 1}, &ListNode{Val: 2}},
	{
		&ListNode{2, &ListNode{Val: 3}},
		&ListNode{5, &ListNode{Val: 3}},
		&ListNode{7, &ListNode{Val: 6}},
	},
	{
		&ListNode{2, &ListNode{4, &ListNode{Val: 3}}},
		&ListNode{5, &ListNode{5, &ListNode{Val: 3}}},
		&ListNode{7, &ListNode{9, &ListNode{Val: 6}}},
	},
	{
		&ListNode{2, &ListNode{4, &ListNode{Val: 3}}},
		&ListNode{5, &ListNode{6, &ListNode{Val: 4}}},
		&ListNode{7, &ListNode{0, &ListNode{Val: 8}}},
	},
	{&ListNode{Val: 5}, &ListNode{Val: 5}, &ListNode{0, &ListNode{Val: 1}}},
	{
		&ListNode{1, &ListNode{Val: 8}},
		&ListNode{Val: 0},
		&ListNode{1, &ListNode{Val: 8}},
	},
}

func testAddTwoNumbers(t *testing.T, addTwoNumbers func(*ListNode, *ListNode) *ListNode) {
	failure := func(tc addTwoNumbersTestCase, got *ListNode) {
		t.Errorf("addTwoNumbers(%v, %v) returned %v, expected %v", tc.l1, tc.l2, got, tc.expected)
	}

	for _, tc := range addTwoNumbersTestCases {
		if got := addTwoNumbers(tc.l1, tc.l2); tc.expected == nil {
			if got != nil {
				failure(tc, got)
			}
		} else {
			for curr, exp := got, tc.expected; exp != nil; curr, exp = curr.Next, exp.Next {
				if curr.Val != exp.Val {
					failure(tc, got)
				}
			}
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testAddTwoNumbers(t, addTwoNumbers)
}
