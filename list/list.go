package list

type item struct {
	data interface{}
	next *item
}

type ListIterator struct {
	c *item
}

func (i *ListIterator) Next() *ListIterator {
	i.c = i.c.next
	return i
}

func (i *ListIterator) HasNext() bool {
	if i.c == nil {
		return false
	}
	return true
}

func (i *ListIterator) Item() interface{} {
	return i.c
}

type List struct {
	Start *item
	End   *item
}

func New() *List {
	l := List{}
	return &l
}

func (l *List) Iterator() *ListIterator {
	return &ListIterator{
		c: l.Start,
	}
}

func (l *List) Add(d interface{}) *List {
	i := item{data: d}

	if l.Start == nil {
		l.Start = &i
		l.End = &i
		return l
	}

	l.End.next = &i
	l.End = l.End.next
	return l
}
