package singleton

import "sync"

type Object struct {
}

var once sync.Once

var obj *Object

func GetInstance() *Object {
	once.Do(func() {
		obj = &Object{}
	})
	return obj
}
