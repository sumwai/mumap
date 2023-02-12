package mumap

import "sync"

type Mumap[K comparable, V interface{}] struct {
	mu  sync.Mutex
	Map map[K]V
}

func (m *Mumap[K, V]) Set(k K, v V) *Mumap[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map[k] = v
	return m
}

func (m *Mumap[K, V]) Get(k K) (v V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.Map[k]
	return
}

func New[K comparable, V interface{}]() (m *Mumap[K, V]) {
	m = new(Mumap[K, V])
	m.Map = map[K]V{}
	return
}
