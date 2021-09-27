package hw04lrucache

type Key string

func (k Key) String() string {
	return string(k)
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	flag := false
	if _, ok := c.Get(key); ok {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
		flag = true
	}

	newCacheItem := &cacheItem{
		key: key.String(),
		value: value,
	}

	if c.capacity == c.queue.Len() {
		lastItem := c.queue.Back()
		key := lastItem.Value.(*cacheItem).key
		c.items[Key(key)] = nil
		c.queue.Remove(lastItem)
	}

	listNode := c.queue.PushFront(newCacheItem)
	c.items[key] = listNode

	return flag
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	listNode, ok := c.items[key]

	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(listNode)

	return listNode.Value.(*cacheItem).value, ok
}

func (c *lruCache) Clear()  {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}