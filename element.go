package dll

type Element[T any] struct {
	data T
	next *Element[T]
	prev *Element[T]
	dll  *DoublyLinkedList[T]
}

func NewElement[T any](data T) *Element[T] {
	return &Element[T]{data: data}
}

func (el *Element[T]) Next() *Element[T] {
	return el.next
}

func (el *Element[T]) HasNext() bool {
	return el.next != nil
}

func (el *Element[T]) HasPrev() bool {
	return el.prev != nil
}

func (el *Element[T]) Prev() *Element[T] {
	return el.prev
}

func (el *Element[T]) Value() T {
	return el.data
}

func (el *Element[T]) ReplaceValue(value T) {
	el.data = value
}
