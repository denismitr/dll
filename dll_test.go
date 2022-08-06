package dll_test

import (
	"testing"

	"github.com/denismitr/dll"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDLL_Remove(t *testing.T) {
	t.Run("remove head", func(t *testing.T) {
		l := dll.New[int]()
		head := dll.NewElement(1)
		l.PushTail(head)

		l.PushTail(dll.NewElement(2))
		l.PushTail(dll.NewElement(3))
		l.PushTail(dll.NewElement(4))

		tail := dll.NewElement(5)
		l.PushTail(tail)

		assert.Equal(t, 5, l.Tail().Value())
		assert.Equal(t, 1, l.Head().Value())

		l.Remove(head)

		assert.Equal(t, 4, l.Len())
		assert.Equal(t, 2, l.Head().Value())
	})

	t.Run("remove tail", func(t *testing.T) {
		l := dll.New[int]()
		head := dll.NewElement(1)

		l.PushTail(head)
		l.PushTail(dll.NewElement(2))
		l.PushTail(dll.NewElement(3))
		l.PushTail(dll.NewElement(4))

		tail := dll.NewElement(5)
		l.PushTail(tail)

		l.Remove(tail)

		assert.Equal(t, 4, l.Len())
		assert.Equal(t, 4, l.Tail().Value())
		assert.Equal(t, 1, l.Head().Value())
	})
}

func TestDll_Push(t *testing.T) {
	t.Run("push remove and check", func(t *testing.T) {
		l := dll.New[string]()

		foo := dll.NewElement("foo")
		bar := dll.NewElement("bar")
		baz := dll.NewElement("baz")

		l.PushTail(foo)
		l.PushTail(bar)

		assert.Equal(t, 2, l.Len())
		assert.True(t, l.Remove(bar))

		l.PushHead(baz)

		assert.Equal(t, "baz", l.Head().Value())
		assert.Equal(t, "foo", l.Tail().Value())
	})
}

func TestDll_Sort(t *testing.T) {
	t.Run("simple reverse", func(t *testing.T) {
		l := dll.New[int]()
		l.PushTail(dll.NewElement(4))
		l.PushTail(dll.NewElement(3))
		l.PushTail(dll.NewElement(2))
		l.PushTail(dll.NewElement(5))
		l.PushTail(dll.NewElement(1))

		comparator := func(a int, b int) bool { return a < b }
		l.Sort(comparator)

		require.NotNil(t, l.Head())
		require.NotNil(t, l.Tail())
		require.NotNil(t, l.Head().Next())
		require.NotNil(t, l.Head().Next().Next())
		require.NotNil(t, l.Head().Next().Next().Next())
		require.NotNil(t, l.Tail().Prev())
		require.NotNil(t, l.Tail().Prev().Prev())

		require.Equal(t, 1, l.Head().Value())
		require.Equal(t, 2, l.Head().Next().Value())
		require.Equal(t, 3, l.Head().Next().Next().Value())
		require.Equal(t, 5, l.Tail().Value())
		require.Equal(t, 4, l.Tail().Prev().Value())
	})

	t.Run("reverse long sequence of integers", func(t *testing.T) {
		l := dll.New[int]()
		for i := 100_000; i > 0; i-- {
			l.PushHead(dll.NewElement(i))
		}

		// simple check befor sort
		assert.Equal(t, 1, l.Head().Value())
		assert.Equal(t, 100_000, l.Tail().Value())

		l.Sort(func(a int, b int) bool { return a > b })

		// simple check after sort
		assert.Equal(t, 100_000, l.Head().Value())
		assert.Equal(t, 1, l.Tail().Value())

		curr := l.Head()
		for i := 100_000; i > 0; i-- {
			require.NotNil(t, curr)
			assert.Equal(t, i, curr.Value())
			curr = curr.Next()
		}
	})
}

func TestDLL_Reverse(t *testing.T) {
	t.Run("reverse sequence of integers", func(t *testing.T) {
		l := dll.New[int]()
		for i := 0; i < 1_000; i++ {
			l.PushTail(dll.NewElement(i))
		}

		// befor reverse
		assert.Equal(t, 0, l.Head().Value())
		assert.Equal(t, 999, l.Tail().Value())

		l.Reverse()

		// after reverse
		assert.Equal(t, 999, l.Head().Value())
		assert.Equal(t, 0, l.Tail().Value())

		// check each element from head to tail after reverse
		curr := l.Head()
		for i := 999; i >= 0; i-- {
			require.NotNil(t, curr)
			assert.Equal(t, i, curr.Value())
			curr = curr.Next()
		}

		// check each element from tail to head after reverse
		curr = l.Tail()
		for i := 0; i < 1_000; i++ {
			require.NotNil(t, curr)
			assert.Equal(t, i, curr.Value())
			curr = curr.Prev()
		}
	})
}
