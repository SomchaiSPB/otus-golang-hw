package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List // Remove me after realization.
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
	}

	if l.head == nil {
		l.head = newListItem
		l.tail = newListItem
	} else {
		current := l.head
		for current.Prev != nil {
			current = current.Prev
		}
		newListItem.Next = current
		current.Prev = newListItem
		l.head = newListItem
	}

	l.len++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
	}

	if l.head == nil {
		l.head = newListItem
		l.tail = newListItem
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		newListItem.Prev = current
		current.Next = newListItem
		newListItem.Next = nil
		l.tail = newListItem
	}

	l.len++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	if l.len == 0 {
		return
	}

	current := l.head
	for current.Next != nil {
		if current == i {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}
		current = current.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.len == 0 || l.head.Value == i.Value {
		return
	}
	if l.Back() == i {
		i.Prev.Next = nil
		l.tail = i.Prev
		i.Prev = nil
		i.Next = l.head
		l.head = i
		return
	}

	current := l.head.Next

	for current.Next != nil {
		if current == i {
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
			i.Prev = nil
			i.Next = l.head
			l.head = i
		}
		current = current.Next
	}
}
