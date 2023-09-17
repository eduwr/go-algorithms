package linkedstack

type node struct {
	item string
	next *node
}

type LinkedStackOfStrings struct {
	first *node
}

func (l *LinkedStackOfStrings) IsEmpty() bool {
	return l.first == nil
}

func (l *LinkedStackOfStrings) Push(s string) {
	old := l.first
	l.first = &node{
		item: s,
		next: old,
	}
}

func (l *LinkedStackOfStrings) Pop() string {
	item := l.first.item
	l.first = l.first.next
	return item
}
