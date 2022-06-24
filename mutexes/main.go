package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
}

// Mutex là viết tắt của mutual exclusion. Có nghĩa là loại trừ lẫn nhau
// kiểu như 1 luồng có 2 thứ đều chạy cùng lúc thì dùng cái này thì sẽ phải chạy từng cái 1
//Golang team khuyên rằng "Don't communicate by sharing memory; share memory by communicating" ~ "Đừng trao đổi bằng cách chia sẻ vùng nhớ; chia sẻ vùng nhớ bằng trao đổi". Ý nói các anh em đồng đạo hãy dùng channel để trao đổi giữa các tác vụ chạy đồng thời, chớ có dùng biến toàn cục. Tôi hiểu là nếu sử dụng channel, thì sẽ tránh được tranh chấp, khóa chết (dead lock).

// https://techmaster.vn/posts/34830/kinh-nghiem-lam-viec-voi-mutex-trong-golang
