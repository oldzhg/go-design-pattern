package singleton

import "sync"

var (
	lazySingleton *Singleton
	lock          sync.Mutex
)

func GetLazySingleton() *Singleton {
	if lazySingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if lazySingleton == nil {
			lazySingleton = &Singleton{}
		}
	}
	return lazySingleton
}
