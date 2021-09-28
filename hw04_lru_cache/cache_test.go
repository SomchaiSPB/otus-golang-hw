package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("key1", 1)
		c.Set("key2", 2)
		c.Set("key3", 3)
		c.Set("key4", 4)

		res, ok := c.Get("key1")

		require.False(t, ok)
		require.Nil(t, res)

		res, ok = c.Get("key2")

		require.True(t, ok)
		require.Equal(t, 2, res)
	})

	t.Run("LRU evict logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("key1", 1)
		c.Set("key2", 2)
		c.Set("key3", 3)

		c.Get("key1")
		c.Get("key2")
		c.Set("key1", 100)
		c.Set("key2", 200)

		res, ok := c.Get("key1")
		require.True(t, ok)
		require.Equal(t, 100, res)

		res, ok = c.Get("key2")
		require.True(t, ok)
		require.Equal(t, 200, res)

		c.Set("key4", 4)

		res, ok = c.Get("key3")
		require.False(t, ok)
		require.Nil(t, res)
	})
}

func TestCacheMultithreading(t *testing.T) {
	//t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

func TestCanSetAndGetCacheItems(t *testing.T) {
	c := NewCache(5)

	c.Set("test_key", 122)
	c.Set("test_key_2", 900)

	res, ok := c.Get("test_key")
	require.True(t, ok)
	require.Equal(t, 122, res)

	res, ok = c.Get("test_key_2")
	require.True(t, ok)
	require.Equal(t, 900, res)
}