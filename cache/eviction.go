package cache

// evict removes an element from the cache based on the eviction strategy.
func (c *Cache) evict() {
	switch c.evictionType {
	case FILO:
		c.evictOldest()
	case LRU:
		c.evictLeastRecentlyUsed()
	}
}

// evictOldest removes the oldest element from the cache (FILO strategy).
func (c *Cache) evictOldest() {
	elem := c.ll.Back()
	if elem != nil {
		c.ll.Remove(elem)
		delete(c.data, elem.Value.(*cacheObject).key)
	}
}

// evictLeastRecentlyUsed removes the least recently used element from the cache (LRU strategy).
func (c *Cache) evictLeastRecentlyUsed() {
	elem := c.ll.Back()
	if elem != nil {
		c.ll.Remove(elem)
		delete(c.data, elem.Value.(*cacheObject).key)
	}
}
