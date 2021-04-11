package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	for x := range Fib(10) {
		fmt.Println(x)
	}

	e := Event{
		Data: 1,
	}
	ob1 := eventObserver{
		ID:   1,
		Time: time.Now(),
	}
	ob2 := eventObserver{
		ID:   2,
		Time: time.Now(),
	}
	sub := eventSubject{
		Observers: make(map[int]Observer),
	}

	sub.AddObserver(ob1)
	sub.AddObserver(ob2)
	sub.Notify(e)
	sub.RemoveObserver(ob1)
	sub.Notify(e)
	sub.AddObserver(ob2)
}
