package putil

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyLock(t *testing.T) {
	k := NewKeyLock[string]()
	m := []string{}
	l := sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	k.Lock("1")

	go func() {
		k.Lock("1")
		defer k.Unlock("1")

		l.Lock()
		m = append(m, "a")
		l.Unlock()

		wg.Done()
	}()

	go func() {
		k.Lock("2")
		defer k.Unlock("2")

		l.Lock()
		m = append(m, "c")
		l.Unlock()

		k.Unlock("1")
		wg.Done()
	}()

	l.Lock()
	m = append(m, "b")
	l.Unlock()
	wg.Wait()

	assert.Equal(t, []string{"b", "c", "a"}, m)
}
