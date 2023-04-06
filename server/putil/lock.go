package putil

import (
	"sync"

	"github.com/reearth/reearthx/util"
)

type KeyLock[T comparable] struct {
	m *util.SyncMap[T, *sync.Mutex]
}

func NewKeyLock[T comparable]() *KeyLock[T] {
	return &KeyLock[T]{
		m: util.NewSyncMap[T, *sync.Mutex](),
	}
}

func (l *KeyLock[T]) Lock(key T) {
	l.getLock(key).Lock()
}

func (l *KeyLock[T]) Unlock(key T) {
	l.getLock(key).Unlock()
}

func (l *KeyLock[T]) getLock(key T) *sync.Mutex {
	nlock := &sync.Mutex{}
	lock, ok := l.m.LoadOrStore(key, nlock)
	if !ok {
		lock = nlock
	}
	return lock
}
