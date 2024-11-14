package routines

import (
	"testing"
	"time"
)

func TestSpawn(t *testing.T) {
	in := newNumGenerator(1, 20)
	out := spawn(square, spawn(filterOdd, in))
	for v := range out {
		println(v)
	}
}

func TestSpawnGroup(t *testing.T) {
	in := newNumGenerator(1, 20)
	out := spawnGroup("square", 2, square, spawnGroup("filterOdd", 3, filterOdd, in))

	time.Sleep(3 * time.Second)
	for v := range out {
		println(v)
	}
}

func newNumGenerator(start, count int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func filterOdd(in int) (int, bool) {
	if in%2 == 0 {
		return 0, false
	}
	return in, true
}

func square(in int) (int, bool) {
	return in * in, true
}
