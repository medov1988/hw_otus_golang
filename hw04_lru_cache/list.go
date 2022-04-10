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
	Length    int
	FrontItem *ListItem
	BackItem  *ListItem
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FrontItem
}

func (l *list) Back() *ListItem {
	return l.BackItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}

	if l.FrontItem != nil {
		l.FrontItem.Prev = newItem
	}

	if l.BackItem == nil {
		l.BackItem = newItem
	}

	newItem.Next = l.FrontItem
	l.FrontItem = newItem
	l.Length++

	return l.FrontItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}

	if l.BackItem != nil {
		l.BackItem.Next = newItem
	}

	if l.FrontItem == nil {
		l.FrontItem = newItem
	}

	newItem.Prev = l.BackItem
	l.BackItem = newItem
	l.Length++

	return l.BackItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.BackItem = i.Prev
	}

	if i.Prev == nil {
		l.FrontItem = i.Next
	}

	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	} else {
		i.Prev.Next = nil
		l.BackItem = i.Prev
	}
	l.Length--

	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
