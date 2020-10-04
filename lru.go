package lru

type LRU interface {
	Len() int
	Cap() int
	Get(key string) interface{}
	Put(key string, value interface{})
}

type cache struct {
	cap  int
	list *list
	kv   map[string]*node
}

func NewLRU(cap int) LRU {
	return &cache{
		cap:  cap,
		list: &list{},
		kv:   make(map[string]*node),
	}
}

func (c *cache) Get(key string) interface{} {
	node, ok := c.kv[key]
	if !ok {
		return nil
	}
	c.list.remove(node)
	c.list.addHead(node.key, node.value)
	return node.value
}

func (c *cache) Put(key string, value interface{}) {
	node, ok := c.kv[key]
	if ok {
		c.list.remove(node)
		c.list.addHead(key, value)
		node.value = value
		return
	}
	if c.isFull() {
		n := c.list.removeTail()
		if n != nil {
			delete(c.kv, n.key)
		}
	}
	node = c.list.addHead(key, value)
	c.kv[key] = node
}

func (c *cache) Len() int {
	return c.list.length
}

func (c *cache) Cap() int {
	return c.cap
}

func (c *cache) isFull() bool {
	return c.list.length == c.cap
}
