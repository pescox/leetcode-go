package main

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead 操作返回 -1 )

type cqueue struct {
	stack1 stack
	stack2 stack
}

// 模拟stack
type stack struct {
	items []int
	len   int
}

func (s *stack) push(val int) {

}

func (s *stack) pop() int {
	return 0
}

func (s *stack) isEmpty() bool {
	return true
}

func NewCQueue() *cqueue {
	return &cqueue{}
}

func (q *cqueue) appendTail(val int) {
	q.stack1.push(val)
}

func (q *cqueue) deleteHead() int {
	// 如果stack2为空，则将stack1中的数据导入到stack2
	if q.stack2.isEmpty() {
		for !q.stack1.isEmpty() {
			q.stack2.push(q.stack1.pop())
		}
	}

	// 如果stack2经过上一轮导入，仍然为空，则说明队列中没有数据，返回-1
	if q.stack2.isEmpty() {
		return -1
	}
	return q.stack2.pop()
}
