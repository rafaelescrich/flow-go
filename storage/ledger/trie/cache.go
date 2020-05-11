package trie

import (
	"container/list"
	"errors"
)

type treeCache interface {
	Add(r Root, t *tree) (evicted bool)
	Get(r Root) (t *tree, ok bool)
	Purge()
	Contains(r Root) bool
}

type lruTreeCache struct {
	size      int
	evictList *list.List
	items     map[string]*list.Element
}

type entry struct {
	key   string
	value *tree
}

func newLRUTreeCache(size int) (*lruTreeCache, error) {
	if size <= 0 {
		return nil, errors.New("must provide a positive size")
	}

	c := &lruTreeCache{
		size:      size,
		evictList: list.New(),
		items:     make(map[string]*list.Element),
	}

	return c, nil
}

func (c *lruTreeCache) Add(r Root, t *tree) (evicted bool) {
	key := keyForRoot(r)

	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = t
		return false
	}

	// Add new item
	ent := &entry{key, t}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size

	// Verify size not exceeded
	if evict {
		return c.removeOldest()
	}

	return
}

func (c *lruTreeCache) Get(r Root) (t *tree, ok bool) {
	key := keyForRoot(r)

	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, true
	}

	return
}

func (c *lruTreeCache) Purge() {
	for _, e := range c.items {
		c.removeElement(e)
	}

	if len(c.items) == 0 {
		c.evictList.Init()
	}
}

func (c *lruTreeCache) Contains(r Root) (ok bool) {
	key := keyForRoot(r)
	_, ok = c.items[key]
	return ok
}

// removeOldest removes the oldest item from the cache.
func (c *lruTreeCache) removeOldest() (evicted bool) {
	ent := c.evictList.Back()
	if ent != nil {
		if c.removeElement(ent) {
			return true
		}

		prev := ent.Prev()
		for prev != nil {
			if c.removeElement(ent) {
				return true
			}

			prev = prev.Prev()
		}
	}

	return false
}

// removeElement is used to remove a given list element from the cache
func (c *lruTreeCache) removeElement(e *list.Element) bool {
	kv := e.Value.(*entry)

	if kv.value.IsLocked() {
		// tree is still in use, do not remove
		return false
	}

	c.evictList.Remove(e)
	delete(c.items, kv.key)

	c.onEvict(kv.key, kv.value)

	return true
}

func (c *lruTreeCache) onEvict(key string, t *tree) {
	_, _ = t.database.SafeClose()
}

func keyForRoot(r Root) string {
	return string(r)
}