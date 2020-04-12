package lock

import "sync"

type Lock struct {
	idLockMap map[string]*sync.RWMutex
}

var (
	RoomLock Lock
)

func init() {
	RoomLock = NewLock()
}

func NewLock() Lock {
	return Lock{idLockMap: make(map[string]*sync.RWMutex)}
}

func (l Lock) Add(id string) {
	l.idLockMap[id] = &sync.RWMutex{}
}

func (l Lock) Delete(id string) {
	delete(l.idLockMap, id)
}

func (l Lock) Lock(id string) {
	l.idLockMap[id].Lock()
}

func (l Lock) Unlock(id string) {
	l.idLockMap[id].Unlock()
}

func (l Lock) RLock(id string) {
	l.idLockMap[id].RLock()
}

func (l Lock) RUnlock(id string) {
	l.idLockMap[id].RUnlock()
}

func (l Lock) RLockAll() {
	for id := range l.idLockMap {
		l.RLock(id)
	}
}

func (l Lock) RUnlockAll() {
	for id := range l.idLockMap {
		l.RUnlock(id)
	}
}
