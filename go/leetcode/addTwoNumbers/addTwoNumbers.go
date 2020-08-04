/**
* You are given two non-empty linked lists representing two non-negative
* integers. The digits are stored in reverse order and each of their
* nodes contain a single digit. Add the two numbers and return it as a
* linked list.
*
* You may assume the two numbers do not contain any leading zero, except
* the number 0 itself.
*
* Example:
*
* Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
* Output: 7 -> 0 -> 8
* Explanation: 342 + 465 = 807.
 */
package addTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ch1, ch2 := make(chan int), make(chan int)

	go l1.Digits(ch1)
	go l2.Digits(ch2)

	head := &ListNode{}
	for tail, sum, v1, v2, ok1, ok2 := head, 0, <-ch1, <-ch2, true, true; ; tail = tail.Next {
		if ok1 {
			sum += v1
			v1, ok1 = <-ch1
		}
		if ok2 {
			sum += v2
			v2, ok2 = <-ch2
		}
		tail.Val = sum % 10
		sum /= 10
		if ok1 || ok2 || sum != 0 {
			tail.Next = &ListNode{}
		} else {
			return head
		}
	}
}
