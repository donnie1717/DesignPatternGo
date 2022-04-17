package singleton

import "sync"

type singleton struct {
}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
	return ins
}

var ins2 *singleton
var mu *sync.Mutex

func GetIns2Or() *singleton {
	if ins2 == nil {
		mu.Lock()
		defer mu.Unlock()
		ins = &singleton{}
	}
	return ins
}

var ins3 *singleton
var once sync.Once

func GetIns3Or() *singleton {
	once.Do(func() {
		ins3 = &singleton{}
	})
	return ins3
}
