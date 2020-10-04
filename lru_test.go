package lru

import "testing"

func TestLRU(t *testing.T) {
	lru := NewLRU(3)
	equalVal(t, 3, lru.Cap(), "cap error")
	equalVal(t, 0, lru.Len(), "len error")

	lru.Put("1", 1)
	lru.Put("2", 2)
	lru.Put("3", 3)
	equalVal(t, 3, lru.Cap(), "put error")
	equalVal(t, 3, lru.Len(), "put error")

	v := lru.Get("1")
	equalVal(t, 1, v, "get error")
	equalVal(t, "1", lru.(*cache).list.head.key, "get error")
	equalVal(t, "2", lru.(*cache).list.tail.key, "get error")

	lru.Put("4", 4)
	equalVal(t, 3, lru.Cap(), "put error")
	equalVal(t, 3, lru.Len(), "put error")
	equalVal(t, "4", lru.(*cache).list.head.key, "get error")
	equalVal(t, "3", lru.(*cache).list.tail.key, "get error")
}
