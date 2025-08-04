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
	len   int
	first *ListItem
	last  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{Value: v, Prev: l.last, Next: nil}

	if l.last != nil {
		l.last.Next = &item
	}

	l.last = &item
	if l.len == 0 {
		l.first = &item
	}

	l.len++

	return &item
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{Value: v, Prev: nil, Next: l.first}

	if l.first != nil {
		l.first.Prev = &item
	}

	l.first = &item
	if l.len == 0 {
		l.last = &item
	}

	l.len++

	return &item
}

func (l *list) Remove(i *ListItem) {
	if i == nil || l.len == 0 {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	i.Value = nil
	i = nil

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.len == 0 || i == l.first {
		return
	}

	if i == l.last {
		l.last = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.first != nil {
		l.first.Prev = i
	}

	i.Prev = nil
	i.Next = l.first

	l.first = i
}

func NewList() List {
	return new(list)
}
