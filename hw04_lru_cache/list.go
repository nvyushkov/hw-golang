package hw04lrucache

type List interface {
	Len() int                          // Длина списка
	Front() *ListItem                  // Первый элемент списка
	Back() *ListItem                   // Последний элемент списка
	PushFront(v interface{}) *ListItem // Добавить значение в начало
	PushBack(v interface{}) *ListItem  // Добавить значение в конец
	Remove(i *ListItem)                // Удалить элемент
	MoveToFront(i *ListItem)           // Переместить элемент в начало
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func (l list) Len() int { return l.len }

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.first
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	if l.len == 0 {
		l.first = item
		l.last = item
	} else {
		item.Next = l.first
		l.first.Prev = item
		l.first = item
	}

	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	item.Prev = l.last
	l.last.Next = item
	l.last = item
	l.len++
	return item
}

func (l *list) Remove(el *ListItem) {
	switch el {
	case l.Front():
		el.Next.Prev = nil
		l.first = el.Next
	case l.Back():
		el.Prev.Next = nil
		l.last = el.Prev
	default:
		el.Prev.Next = el.Next
		el.Next.Prev = el.Prev
	}
	el.Prev = nil
	el.Next = nil
	el.Value = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.first == i {
		return
	}
	if l.last == i {
		l.last = i.Prev
		l.last.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Next = l.first
	i.Prev = nil
	l.first.Prev = i
	l.first = i
}

func NewList() List {
	return new(list)
}
