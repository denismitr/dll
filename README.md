# Doubly linked list implementation for GO with generics

#### Push

You can push from head or from tail
```go
l := dll.New[string]()

foo := dll.NewElement("foo")
bar := dll.NewElement("bar")
baz := dll.NewElement("baz")

l.PushTail(foo)
l.PushTail(bar)

l.Len() // 2

l.Remove(foo) // true

l.PushHead(baz)

l.Head().Value() // "baz"
l.Tail().Value() // "bar"
```

#### Sort

You can sort using your own comparator function
```go
// create your own less func of that type:
type LessFn[T any] func(a T, b T) (less bool)
// like so

l := dll.New[int]()

l.PushTail(dll.NewElement(4))
l.PushTail(dll.NewElement(3))
// ...

lessFn := func(a int, b int) bool { return a < b }
l.Sort(lessFn)

l.Head().Value() // 3
l.Tail().Value() // 4
```

### Characteristics
- Sorting with merge sort O(n*log n)
- Remove O(1)
- PushHead and PushTail O(1)
- Reverse O(n)
- Insert with preserving the sort order O(n) when already sorted and O(n*log n) when not yet sorted