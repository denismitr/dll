# Doubly linked list implementation for GO with generics

```go
l := dll.New[string]()

foo := &dll.NewElement("foo")
bar := &dll.NewElement("bar")
baz := &dll.NewElement("baz")

l.PushTail(foo)
l.PushTail(bar)

l.Len() // 2

l.Remove(foo) // true

l.PushHead(baz)

l.Head() // baz
l.Tail() // foo
```

### Extra Features
- Sorting with merge sort O(n log n)
- Remove O(1)
- Add O(1)