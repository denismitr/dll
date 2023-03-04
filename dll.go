package dll

import "errors"

var (
	ErrNoLessFuncSet     = errors.New("sorting is impossible: no less func set: hint call SortAndPreserveOrderOnInsert")
	ErrInconsistentState = errors.New("doubly linked list is in inconsistent state")
)

type DoublyLinkedList[T any] struct {
	head     *Element[T]
	tail     *Element[T]
	elements int
	lessFn   LessFn[T]
}

// New doubly linked list
func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Reverese the doubly linked list
func (l *DoublyLinkedList[T]) Reverse() {
	var temp *Element[T]
	curr := l.head
	tailTemp := l.head

	for curr != nil {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}

	if temp != nil {
		l.head = temp.prev
	}

	l.tail = tailTemp
}

// Remove element from doubly linked list and set all its's links to nil
func (l *DoublyLinkedList[T]) Remove(el *Element[T]) bool {
	if el.dll != l {
		return false
	}

	// is head
	if el.prev == nil {
		l.head = el.next
	} else if el.next != nil {
		el.next.prev = el.prev
	}

	// is tail
	if el.next == nil {
		l.tail = el.prev
	}

	if el.prev != nil {
		el.prev.next = el.next
	}

	el.next = nil
	el.prev = nil
	el.dll = nil

	l.elements--
	return true
}

// Len returns the length of the doubly linked list
func (l *DoublyLinkedList[T]) Len() int {
	return l.elements
}

// PushHead - pushes element to the head of the doubly linked list
func (l *DoublyLinkedList[T]) PushHead(el *Element[T]) {
	// ensure correctness
	el.next = nil
	el.prev = nil

	l.elements++
	el.dll = l

	if l.head == nil {
		l.head = el
		l.tail = el
	} else {
		oldHead := l.head
		el.next = oldHead
		oldHead.prev = el
		l.head = el
	}
}

// PushHead - pushes element to the tail of the doubly linked list
func (l *DoublyLinkedList[T]) PushTail(el *Element[T]) {
	l.elements++
	el.dll = l

	// ensure correctness
	el.next = nil
	el.prev = nil

	if l.head == nil {
		l.head = el
		l.tail = el
	} else {
		prev := l.tail
		prev.next = el
		el.prev = prev
		l.tail = el
		el.next = nil
	}
}

// Head - returns the head of the doubly linked list
func (l *DoublyLinkedList[T]) Head() *Element[T] {
	return l.head
}

// Tail - returns the tail of the doubly linked list
func (l *DoublyLinkedList[T]) Tail() *Element[T] {
	return l.tail
}

// LessFn - compares a and b of type T and returns boolean
// indicating weather a is less than b
type LessFn[T any] func(a T, b T) (less bool)

// Sort the doubly linked list using the comparator function
func (l *DoublyLinkedList[T]) Sort(lessFn LessFn[T]) {
	l.head, l.tail = mergeSort(l.head, l.tail, lessFn)
}

func mergeSort[T any](
	head *Element[T],
	tail *Element[T],
	comparator LessFn[T],
) (newHead *Element[T], newTail *Element[T]) {
	if head == nil || head == tail || head.next == nil {
		return head, tail
	}

	// find the relative middle node
	middle := middle(head, tail)
	right, rightTail := mergeSort(middle.next, tail, comparator)

	// separating sublist
	if middle.next != nil {
		middle.next.prev = nil
	}
	middle.next = nil

	// get the left sublist
	left, leftTail := mergeSort(head, middle, comparator)

	return merge(left, leftTail, right, rightTail, comparator)
}

// middle finds middle element using the fast/slow pointer strategy
func middle[T any](
	head *Element[T],
	tail *Element[T],
) *Element[T] {
	if head == nil || head == tail {
		return head
	}

	slow := head
	fast := head
	for fast != nil && fast.next != nil && fast.next.next != nil {
		if fast.next == tail || fast.next.next == tail {
			break
		}

		fast = fast.next.next
		slow = slow.next
	}

	// this is middle element
	return slow
}

// merge elements together
func merge[T any](
	left *Element[T],
	leftTail *Element[T],
	right *Element[T],
	rightTail *Element[T],
	comparator LessFn[T],
) (*Element[T], *Element[T]) {
	if left == nil {
		return right, rightTail
	}

	if right == nil {
		return left, leftTail
	}

	// pick the smallest key
	if comparator(left.data, right.data) {
		next, tail := merge(left.next, leftTail, right, rightTail, comparator)
		left.next = next
		left.next.prev = left
		left.prev = nil
		return left, tail
	} else {
		next, tail := merge(left, leftTail, right.next, rightTail, comparator)
		right.next = next
		right.next.prev = right
		right.prev = nil
		return right, tail
	}
}
