package test

import (
	"fmt"
	"testing"
)

func bucketShift(b uint8) uintptr {
	return uintptr(1) << b
}

func overLoadFactor(count int, B uint8) bool {
	return count > 8 && uintptr(count) > 13*(bucketShift(B)/2)
}

func bucketLen(hint int) uint8 {
	B := uint8(0)
	for overLoadFactor(hint, B) {
		B++
	}
	return B
}

func TestMapB(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, bucketLen(i))
	}
}

/*
(1) size <9，bucketLen = 1 (2^0) 		B=0
(2) size < 14，bucketLen = 2 (2^1) 		B=1
(3) size < 27，bucketLen = 4 (2^2) 		B=2
(4) size < 53，bucketLen = 8 (2^3) 		B=3
(5) size < 104，bucketLen = 16 (2^4) 	B=4


第二行与第一行size之差A=5，bucket长度差为1，A/B=5。
第三行与第二行size之差A=13，bucket长度差B=2，A/B=6.5。
第四行与第三行size之差A=26，bucket长度差B=4，A/B=6.5。
第五行与第四行size之差A=51，bucket长度差B=8，A/B=6.375

*/