package dll

type Element[T any] struct {
	Data T
	next *Element[T]
	prev *Element[T]
	dll  *DoublyLinkedList[T]
}

func NewElement[T any](data T) *Element[T] {
	return &Element[T]{Data: data}
}

func (el *Element[T]) Next() *Element[T] {
	return el.next
}

func (el *Element[T]) Prev() *Element[T] {
	return el.prev
}

func (el *Element[T]) Value() T {
	return el.Data
}
