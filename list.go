package lru

import (
	"fmt"
	"strings"
)

type node struct {
	key        string
	value      interface{}
	prev, next *node
}

type list struct {
	length int
	head   *node
	tail   *node
}

func (l *list) isEmpty() bool {
	return l.length == 0
}

func (l *list) addHead(key string, value interface{}) *node {
	node := &node{
		key:   key,
		value: value,
	}
	defer func() {
		l.length++
	}()
	if l.isEmpty() {
		l.head = node
		l.tail = l.head
		return node
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
	return node
}

func (l *list) removeTail() *node {
	if l.isEmpty() {
		return nil
	}

	res := l.tail
	if res == nil {
		return nil
	}

	l.length--

	// if only one item
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return res
	}

	l.tail = res.prev
	if l.tail != nil {
		l.tail.next = nil
	}

	return res
}

func (l *list) remove(n *node) {
	if n == nil {
		return
	}
	if l.isEmpty() {
		return
	}

	l.length--

	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
}

func (l *list) String() string {
	var res string
	for n := l.head; n != nil; n = n.next {
		res += fmt.Sprintf("%+v -> ", n)
	}
	return strings.TrimSuffix(res, " -> ")
}

func (n *node) String() string {
	return fmt.Sprintf("(key: %s, val: %+v)", n.key, n.value)
}
