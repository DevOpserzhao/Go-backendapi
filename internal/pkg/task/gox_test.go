package task

import (
	"log"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	ch := make(chan int)
	data := []int{1, 2, 3, 4, 5}

	Go(func() {
		log.Println("In data: start")
		for _, v := range data {
			ch <- v
			time.Sleep(time.Microsecond)
			panic(1)
		}
		log.Println("In data: end")
	})

	t.Helper()

	Go(func() {
		for v := range ch {
			log.Printf("Out data: %v", v)
		}
	})

	time.Sleep(time.Second)
}

func TestGo2(t *testing.T) {
	ch := make(chan int)
	data := []int{1, 2, 3, 4, 5}

	go func() {
		log.Println("In data: start")
		for _, v := range data {
			ch <- v
			time.Sleep(time.Microsecond)
			// panic("某些原因goroutine panic")
		}
		log.Println("In data: end")
	}()

	go func() {
		for v := range ch {
			log.Printf("Out data: %v", v)
		}
	}()

	time.Sleep(time.Second)
}