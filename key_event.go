package microbot

import (
	"container/list"
	"fmt"
	"sync"
)

type KeyEventList struct {
	data *list.List
	max  int
}

func NewQueue(max int) *KeyEventList {
	q := new(KeyEventList)
	q.data = list.New()
	q.max = max
	return q
}

func NewKeyEvent(typ string, content string) {

}

func (q *KeyEventList) push(v interface{}) {
	if q.data.Len() >= q.max {
		q.pop()
	}
	defer lock.Unlock()
	lock.Lock()
	q.data.PushFront(v)
}

func (q *KeyEventList) pop() interface{} {
	defer lock.Unlock()
	lock.Lock()
	iter := q.data.Back()
	v := iter.Value
	q.data.Remove(iter)
	return v
}

func (q *KeyEventList) dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}

var lock sync.RWMutex
