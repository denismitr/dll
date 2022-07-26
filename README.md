# Doubly linked list implementation for GO with generics

```go
l := dll.New[string]()

foo := &dll.Elemenet[string]{Data: "foo"}
bar := &dll.Elemenet[string]{Data: "bar"}

l.PushTail(foo)
l.PushTail(bar)

l.Len() // 2

l.Remove(foo) // true
```

### Extra Features
- Sorting with merge sort
- Remove O(1)
- Add O(1)