package cache

import "time"

type elem struct {
	value    string
	deadline time.Time
}

type Cache struct {
	elems map[string]elem
	delay time.Duration
}

func NewCache() Cache {
	return Cache{elems: make(map[string]elem), delay: 55 * time.Millisecond}
}

func (c *Cache) Get(key string) (string, bool) {
	e, ok := c.elems[key]
	if !ok {
		return "", ok
	}
	return e.value, true
}

func (c *Cache) Put(key, value string) {
	c.elems[key] = elem{value: value, deadline: time.Now()}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0)
	for key, _ := range c.elems {
		keys = append(keys, key)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.elems[key] = elem{value: value, deadline: deadline}
	go c.expire(key)
}

func (c *Cache) expire(key string) {
	for c.elems[key].deadline.After(time.Now()) {
		time.Sleep(c.delay)
	}
	delete(c.elems, key)
}
