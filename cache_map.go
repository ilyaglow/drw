package drw

import (
	"sync"
)

// MapCache is a simple Go map.
type MapCache struct {
	sync.RWMutex
	cache map[string]struct{}
}

// NewMapCache initializes the MapCache.
func NewMapCache(size int) *MapCache {
	if size != 0 {
		return &MapCache{
			cache: make(map[string]struct{}, size),
		}
	}
	return &MapCache{
		cache: make(map[string]struct{}),
	}
}

// Set implements cache interface.
func (m *MapCache) Set(b []byte) (exists bool, err error) {
	sb := string(b)
	// log.Printf("storing %s", sb)
	m.Lock()
	defer m.Unlock()
	_, ok := m.cache[sb]
	if ok {
		return true, nil
	}

	m.cache[sb] = struct{}{}
	return false, nil
}
