package singleton

import "sync"

type singleton struct {
	Name string
}

var (
	ins  *singleton
	once sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		ins = &singleton{
			Name: "jack",
		}
	})
	return ins
}
