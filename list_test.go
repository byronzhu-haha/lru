package lru

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	l := &list{}
	l.addHead("1", 1)
	equalVal(t, 1, l.length, "addHead error")
	equalVal(t, "1", l.head.key, "addHead error")
	equalVal(t, "1", l.tail.key, "addHead error")
	l.addHead("2", 2)
	equalVal(t, 2, l.length, "addHead error")
	equalVal(t, "2", l.head.key, "addHead error")
	equalVal(t, "1", l.tail.key, "addHead error")

	l.addHead("3", 3)

	t.Log(l)

	tail := l.removeTail()
	equalVal(t, "1", tail.key, "removeTail error")
	equalVal(t, 2, l.length, "removeTail error")
	equalVal(t, "2", l.tail.key, "removeTail error")
	equalVal(t, "3", l.head.key, "removeTail error")

	t.Log(l)

	tail = l.removeTail()
	equalVal(t, "2", tail.key, "removeTail error")
	equalVal(t, 1, l.length, "removeTail error")
	equalVal(t, "3", l.head.key, "removeTail error")
	equalVal(t, "3", l.tail.key, "removeTail error")

	t.Log(l)

	tail = l.removeTail()
	equalVal(t, "3", tail.key, "removeTail error")
	equalVal(t, 0, l.length, "removeTail error")
	equalNode(t, nil, l.head, "removeTail error")
	equalNode(t, nil, l.tail, "removeTail error")

	t.Log(l)

	n4 := l.addHead("4", 4)
	n5 := l.addHead("5", 5)
	n6 := l.addHead("6", 6)

	t.Log(l)

	l.remove(n5)
	equalVal(t, 2, l.length, "remove error")
	equalNode(t, n6, l.head, "remove error")
	equalNode(t, n4, l.tail, "remove error")

	l.remove(n6)
	equalVal(t, 1, l.length, "remove error")
	equalNode(t, n4, l.head, "remove error")
	equalNode(t, n4, l.tail, "remove error")

	l.remove(n4)
	equalVal(t, 0, l.length, "remove error")
	equalNode(t, nil, l.head, "remove error")
	equalNode(t, nil, l.tail, "remove error")
}

func equalVal(t *testing.T, expect, val interface{}, errWithMsg string) {
	if reflect.DeepEqual(expect, val) {
		return
	}
	printf(t, expect, val, errWithMsg)
}

func equalNode(t *testing.T, expect, val *node, errWithMsg string)  {
	if expect == val {
		return
	}
	if expect == nil || val == nil {
		printf(t, expect, val, errWithMsg)
		return
	}

	if expect.key == val.key && expect.value == val.value {
		return
	}

	printf(t, expect, val, errWithMsg)
}

func printf(t *testing.T, expect, val interface{}, errWithMsg string) {
	_, file, line, _ := runtime.Caller(2)
	gopath := os.Getenv("GOPATH") + "/src/"
	if idx := strings.Index(file, gopath); idx != -1 {
		file = file[len(gopath):]
	}
	t.Error(fmt.Sprintf("[%s:%d] %s, val(%+v) should be equal to expect(%+v)", file, line, errWithMsg, val, expect))
}