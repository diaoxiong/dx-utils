package routines

import (
	"log"
	"runtime/debug"
)

func SafeGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("恢复panic: %v\n堆栈:\n%s", r, debug.Stack())
			}
		}()
		f()
	}()
}
