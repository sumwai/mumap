package mumap

import (
	"strconv"
	"sync"
	"testing"
)

func TestMumap(t *testing.T) {
	type Point struct {
		X, Y int
	}
	var c int = 1e7 + 1

	m := New[string, *Point]()
	var i int
	wg := sync.WaitGroup{}
	wg.Add(c + 1)
	for i = 0; i <= c; i++ {
		go func(i int) {
			m.Set(strconv.Itoa(i), &Point{i, i})
			wg.Done()
		}(i)
	}
	wg.Wait()

	if d, ok := m.Get(strconv.Itoa(c)); !ok || d.X != c || d.Y != c {
		t.Fatal("Error", d, ok)
	}
}
