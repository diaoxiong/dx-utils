package routines

import (
	"fmt"
	"sync"
)

func spawn(f func(int) (int, bool), in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			r, ok := f(v)
			if ok {
				out <- r
			}
		}
		close(out)
	}()
	return out
}

func spawnGroup(name string, num int, f func(int) (int, bool), in <-chan int) <-chan int {
	groupOut := make(chan int)
	var outSlice []chan int
	for i := 0; i < num; i++ {
		out := make(chan int)
		go func(i int) {
			name := fmt.Sprintf("%s-%d", name, i)
			fmt.Printf("%s begin to work...\n", name)

			for v := range in {
				r, ok := f(v)
				if ok {
					out <- r
				}
			}
			close(out)
			fmt.Printf("%s work done\n", name)
		}(i)
		outSlice = append(outSlice, out)
	}

	go func() {
		var wg sync.WaitGroup
		for _, out := range outSlice {
			wg.Add(1)
			go func(out <-chan int) {
				for v := range out {
					groupOut <- v
				}
				wg.Done()
			}(out)
		}
		wg.Wait()
		close(groupOut)
	}()

	return groupOut
}
