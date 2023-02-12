package mumap

import "sync"

// mumap struct
type Mumap[K comparable, V interface{}] struct {
	mu  sync.Mutex
	Map map[K]V
}

// EachFunc if return true, that will be break, otherwise, continue
type EachFunc[K comparable, V interface{}] func(K, V) bool

// Set like `map[k] = v`, but return mumap
func (m *Mumap[K, V]) Set(k K, v V) *Mumap[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map[k] = v
	return m
}

// Get returns value and ok, like `v, ok := map[k]`
func (m *Mumap[K, V]) Get(k K) (v V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.Map[k]
	return
}

// Each range the mumap, call that item uses EachFunc.
// that for loop will break if EachFunc returns `true`
// otherwise will be continue
func (m *Mumap[K, V]) Each(do EachFunc[K, V]) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.Map {
		if do(k, v) {
			break
		}
	}
}

// New Create a new mumap
func New[K comparable, V interface{}]() (m *Mumap[K, V]) {
	m = new(Mumap[K, V])
	m.Map = map[K]V{}
	return
}
