package dll

type DoublyLinkedList[T any] struct {
	head     *Element[T]
	tail     *Element[T]
	elements int
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

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
	} else if el.prev != nil {
		el.prev.next = el.next
	}

	el.next = nil
	el.prev = nil
	el.dll = nil

	l.elements--
	return true
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.elements
}

func (l *DoublyLinkedList[T]) PushHead(el *Element[T]) {
	if l.head == nil {
		l.head = el
		l.tail = el
	} else {
		oldHead := l.head
		el.next = oldHead
		oldHead.prev = el
		l.head = el
	}
	l.elements++
	el.dll = l
}

func (l *DoublyLinkedList[T]) PushTail(el *Element[T]) {
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
	l.elements++
	el.dll = l
}

func (l *DoublyLinkedList[T]) Head() *Element[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) Tail() *Element[T] {
	return l.tail
}

type CompareFn[T any] func(T, T) bool

func (l *DoublyLinkedList[T]) Sort(comparator CompareFn[T]) {
	l.head, l.tail = mergeSort(l.head, l.tail, comparator)
}

func mergeSort[T any](
	head *Element[T],
	tail *Element[T],
	comparator CompareFn[T],
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
	comparator CompareFn[T],
) (*Element[T], *Element[T]) {
	if left == nil {
		return right, rightTail
	}

	if right == nil {
		return left, leftTail
	}

	// pick the smallest key
	if comparator(left.Data, right.Data) {
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
