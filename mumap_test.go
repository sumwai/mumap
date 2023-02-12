package mumap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

var count int = 1e1

func TestMumapGetSet(t *testing.T) {
	type Point struct {
		X, Y int
	}
	var c int = count + 1

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

func TestMumapEach(t *testing.T) {
	var c int = count
	var want int = 0
	var got int = 0
	var num int = 100

	m := New[int, int]()

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(c + 1)
	// Set Value
	for i := 0; i <= c; i++ {
		go func(i int) {
			n := rand.Intn(1000)
			// want++, if that random number is `num`
			if n == num {
				mutex.Lock()
				want++
				mutex.Unlock()
			}
			m.Set(i, n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	m.Each(func(_, v int) bool {
		if v == num {
			got++
		}
		return false
	})

	if got != want {
		t.Fatalf("Each error, Want: %d, but Got: %d", want, got)
	}
}

func TestFilter(t *testing.T) {
	m := New[int, int]()

	var c int = count

	want := 0
	got := 0

	for i := 0; i < c; i++ {
		if i%100 == 0 {
			want++
		}
		m.Set(i, i)
	}

	m2 := m.Filter(func(k, v int) bool {
		return v%100 == 0
	})

	got = m2.Length()

	if got != want {
		t.Fatalf("Filter error, want: %d, but got: %d", want, got)
	}
}
