package util

import "sync"

type MutexKey string // mutex key used to identify different mutexes

type kMutexItem struct { // key based mutex item
	sync.Mutex
	refCnt uint32
}

// KMutex key based multiple mutex, which is used to help improve locking performance by
// preventing a giant time-consuming locking.
type KMutex struct {
	l sync.Mutex               // operation lock
	s map[MutexKey]*kMutexItem // key mutex mapping
	p *sync.Pool               // mutex pool
}

func NewKMutex() *KMutex {
	return &KMutex{
		s: make(map[MutexKey]*kMutexItem),
		p: &sync.Pool{
			New: func() interface{} {
				return &kMutexItem{refCnt: 0}
			},
		},
	}
}

// Lock locks by key
func (km *KMutex) Lock(key MutexKey) {
	km.l.Lock()

	lock, ok := km.s[key]
	if !ok {
		lock = km.p.Get().(*kMutexItem)
		km.s[key] = lock
	}
	lock.refCnt++

	km.l.Unlock() // must unlock km.l first, otherwise the next lock may block

	lock.Lock()
}

// Unlock unlocks by key
func (km *KMutex) Unlock(key MutexKey) {
	km.l.Lock()
	defer km.l.Unlock()

	lock, ok := km.s[key]
	if !ok || lock.refCnt == 0 {
		panic("must lock mutex before unlock")
	}

	lock.refCnt--
	if lock.refCnt == 0 { // put back to sync pool
		km.p.Put(lock)
		delete(km.s, key)
	}

	lock.Unlock()
}
