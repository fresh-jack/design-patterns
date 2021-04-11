package observer

import (
	"fmt"
	"sync"
	"time"
)

// Event ...
type Event struct {
	Data int
}

// Observer ...
type Observer interface {
	NotifyCallback(e Event)
	GetID() int
}

// Subject ...
type Subject interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	Notify(e Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}

func (e eventObserver) NotifyCallback(data Event) {
	fmt.Println("event: ", data.Data)
}

func (e eventObserver) GetID() int {
	return e.ID
}

type eventSubject struct {
	sync.Mutex
	Observers map[int]Observer
}

func (e eventSubject) AddObserver(o Observer) {
	e.Lock()
	defer e.Unlock()
	if _, ok := e.Observers[o.GetID()]; ok {
		fmt.Println("observer already exist!")
		return
	}
	e.Observers[o.GetID()] = o
}

func (e *eventSubject) RemoveObserver(o Observer) {
	e.Lock()
	defer e.Unlock()
	if _, ok := e.Observers[o.GetID()]; ok {
		delete(e.Observers, o.GetID())
	}
	return
}

func (e *eventSubject) Notify(data Event) {
	for _, v := range e.Observers {
		v.NotifyCallback(data)
	}
}

func Fib(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}
