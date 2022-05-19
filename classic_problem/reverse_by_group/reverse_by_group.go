package reverse_by_group

// 给定一个单链表的头节点 head,实现一个调整单链表的函数。
// 使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。
//（不能使用队列或者栈作为辅助）
// 例如：
// 链表：1->2->3->4->5->6->7->8->null, K = 3。
// 那么 6->7->8，3->4->5，1->2 分别为一组。调整后为：1->2->5->4->3->8->7->6->null。其中1，2不调整，因为不够一组。

type Node struct {
	data int
	next *Node
}

func ReverseList(head *Node) *Node {
	dummy := &Node{}
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = dummy.next
		dummy.next = cur
		cur = tmp
	}
	return dummy.next
}

func ReverseByGroup(head *Node, K int) *Node {
	// 前进K-1步，到达第K个节点
	cur := head
	for i := 0; i < K-1 && cur != nil; i++ {
		cur = cur.next
	}
	// 如果cur为nil，说明未到达第K个节点链表就已经遍历结束，节点数量不够K个，直接返回
	if cur == nil {
		return head
	}
	// 剪断第K个节点与后面的链接，剪断前，先保存后面链表的头节点
	leftHead := cur.next
	cur.next = nil
	// 翻转前K个节点，更新新头节点与新尾节点
	newHead := ReverseList(head)
	newTail := head // 新的尾节点就是翻转前的头节点
	// 递归执行后面链表
	newLeftHead := ReverseByGroup(leftHead, K)
	// 链接两个链表
	newTail.next = newLeftHead
	return newHead
}

func Solve(head *Node, K int) *Node {
	// 先整体翻转链表
	head = ReverseList(head) // 翻转前：1->2->3->4->5->6->7->8->null 翻转后：8->7->6->5->4->3->2->1->null
	// 再分组翻转链表
	head = ReverseByGroup(head, K) // 分组翻转前：8->7->6->5->4->3->2->1->null 分组翻转后：6->7->8->3->4->5->2->1->null
	// 最后又一次整体翻转链表
	head = ReverseList(head) // 翻转前：6->7->8->3->4->5->2->1->null 翻转后：1->2->5->4->3->8->7->6->null
	return head
}
