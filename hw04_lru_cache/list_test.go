package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestCanAddToLinkedListHead(t *testing.T) {
	l := NewList()

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	require.Equal(t, 3, l.Len())
	require.Equal(t, 3, l.Front().Value)
	require.Equal(t, 1, l.Back().Value)
}

func TestCanAddToLinkedListTail(t *testing.T) {
	l := NewList()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)

	require.Equal(t, 4, l.Len())
	require.Equal(t, 10, l.Front().Value)
	require.Equal(t, 40, l.Back().Value)
}

func TestCanRemoveElementFromList(t *testing.T) {
	l := NewList()

	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)

	l.Remove(l.Front().Next)

	require.Equal(t, 3, l.Len())
	require.Equal(t, 10, l.Front().Value)
	require.Equal(t, 30, l.Front().Next.Value)
	require.Equal(t, 40, l.Back().Value)
}

func TestCanMoveItemToFront(t *testing.T) {
	l := NewList()

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)

	l.MoveToFront(l.Back())

	require.Equal(t, 1, l.Front().Value)
	require.Equal(t, 4, l.Front().Next.Value)
	require.Equal(t, 2, l.Back().Value)

	// 1 - 4 - 3 - 2
	// 3 - 1 - 4 - 2
	l.MoveToFront(l.Back().Prev)

	require.Equal(t, 3, l.Front().Value)
	require.Equal(t, 1, l.Front().Next.Value)
	require.Equal(t, 4, l.Back().Prev.Value)
	require.Equal(t, 2, l.Back().Value)

	ll := NewList()

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)
	ll.PushFront(4)

	ll.MoveToFront(ll.Back().Prev)

	require.Equal(t, 2, ll.Front().Value)
	require.Equal(t, 4, ll.Front().Next.Value)
	require.Equal(t, 1, ll.Back().Value)
}