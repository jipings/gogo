package engine

import (
	"testing"
	"time"
)

// 想要验证channel写是否会阻塞，发现test主函数里第一行的输出也被阻塞了
// 这是为什么呢？
func myChan(in chan int, t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Errorf("I' m not blocked")
		in <- 1
	}

}

var ch = time.Tick(time.Millisecond)

func readChanFromOne(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(n int, t *testing.T) {
			<-ch
			t.Logf("No #%d", n)

		}(i, t)
	}

}

func TestChannel(t *testing.T) {
	// t.Errorf("I' m working")
	// in := make(chan int)
	// myChan(in, t)
	timer := time.After(time.Second * 5)
	readChanFromOne(t)
	<-timer
}
