package observable

import (
	"awesomeProject/designPatterns/observer/observer"
	"container/list"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x observer.Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x observer.Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(observer.Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(observer.Observer).Notify(data)
	}
}
