package microbot

import (
	"container/list"
	"fmt"
	"sync"
)

type KeyframeList struct {
	data *list.List
	max  int
}

func NewQueue(max int) *KeyframeList {
	q := new(KeyframeList)
	q.data = list.New()
	q.max = max
	return q
}

func (q *KeyframeList) push(v interface{}) {
	if q.data.Len() >= q.max {
		q.pop()
	}
	defer lock.Unlock()
	lock.Lock()
	q.data.PushFront(v)
}

func (q *KeyframeList) pop() interface{} {
	defer lock.Unlock()
	lock.Lock()
	iter := q.data.Back()
	v := iter.Value
	q.data.Remove(iter)
	return v
}

func (q *KeyframeList) dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}

var lock sync.RWMutex
