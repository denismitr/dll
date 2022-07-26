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
		head := &dll.Element[int]{Data: 1}
		l.PushTail(head)

		l.PushTail(&dll.Element[int]{Data: 2})
		l.PushTail(&dll.Element[int]{Data: 3})
		l.PushTail(&dll.Element[int]{Data: 4})

		tail := &dll.Element[int]{Data: 5}
		l.PushTail(tail)

		assert.Equal(t, 5, l.Tail().Data)
		assert.Equal(t, 1, l.Head().Data)

		l.Remove(head)

		assert.Equal(t, 4, l.Len())
		assert.Equal(t, 2, l.Head().Data)
	})

	t.Run("remove tail", func(t *testing.T) {
		l := dll.New[int]()
		head := &dll.Element[int]{Data: 1}

		l.PushTail(head)
		l.PushTail(&dll.Element[int]{Data: 2})
		l.PushTail(&dll.Element[int]{Data: 3})
		l.PushTail(&dll.Element[int]{Data: 4})

		tail := &dll.Element[int]{Data: 5}
		l.PushTail(tail)

		l.Remove(tail)

		assert.Equal(t, 4, l.Len())
		assert.Equal(t, 4, l.Tail().Data)
		assert.Equal(t, 1, l.Head().Data)
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

		assert.Equal(t, "baz", l.Head().Data)
		assert.Equal(t, "foo", l.Tail().Data)
	})
}

func TestDll_Sort(t *testing.T) {
	t.Run("simple reverse", func(t *testing.T) {
		l := dll.New[int]()
		l.PushTail(&dll.Element[int]{Data: 4})
		l.PushTail(&dll.Element[int]{Data: 3})
		l.PushTail(&dll.Element[int]{Data: 2})
		l.PushTail(&dll.Element[int]{Data: 5})
		l.PushTail(&dll.Element[int]{Data: 1})

		comparator := func(a int, b int) bool { return a < b }
		l.Sort(comparator)

		require.NotNil(t, l.Head())
		require.NotNil(t, l.Tail())
		require.NotNil(t, l.Head().Next())
		require.NotNil(t, l.Head().Next().Next())
		require.NotNil(t, l.Head().Next().Next().Next())
		require.NotNil(t, l.Tail().Prev())
		require.NotNil(t, l.Tail().Prev().Prev())

		require.Equal(t, 1, l.Head().Data)
		require.Equal(t, 2, l.Head().Next().Data)
		require.Equal(t, 3, l.Head().Next().Next().Data)
		require.Equal(t, 5, l.Tail().Data)
		require.Equal(t, 4, l.Tail().Prev().Data)
	})
}
