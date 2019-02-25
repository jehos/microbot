package microbot

import (
	"container/list"
	"net/http"
	"sync"

	"github.com/elvinchan/microbot/utils"
)

var (
	lock         sync.RWMutex
	keyEventList KeyEventList
)

type KeyEvent struct {
	Type    string
	Content string
}

type KeyEventList struct {
	data *list.List
	max  int
}

func KeyEventController() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v []interface{}
		for iter := keyEventList.data.Back(); iter != nil; iter = iter.Prev() {
			v = append(v, iter.Value)
		}
		utils.RenderJson(w, v)
	})
}

func NewQueue(max int) *KeyEventList {
	q := new(KeyEventList)
	q.data = list.New()
	q.max = max
	return q
}

func NewKeyEvent(t string, c string) {
	go keyEventList.push(KeyEvent{
		Type:    t,
		Content: c,
	})
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
