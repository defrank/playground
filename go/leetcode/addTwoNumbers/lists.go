package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Digits(ch chan int) {
	if l == nil {
		ch <- 0
	}
	for ; l != nil; l = l.Next {
		ch <- l.Val
	}
	close(ch)
}
