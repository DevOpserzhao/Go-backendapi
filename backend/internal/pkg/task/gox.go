package task

import "log"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 上报问题
				log.Println(err)
			}
		}()
		f()
	}()
}
