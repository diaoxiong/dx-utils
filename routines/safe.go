package routines

import "log"

func SafeGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("恢复panic: %v", r)
			}
		}()
		f()
	}()
}
