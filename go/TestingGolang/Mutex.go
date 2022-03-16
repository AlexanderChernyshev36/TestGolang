package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	my sync.Mutex
	c map[string]int 
}

func (c *Counter) Inc(key string) {
	c.my.Lock()
c.c[key]++
c.my.Unlock()
}

func (c *Counter) Value(key string)int{
	c.my.Lock()
    c.my.Unlock()
return c.c[key]
}

func main() {
	key := "test"
c := Counter{c: make(map[string]int)}
for i := 0; i < 1000; i++ {
	go c.Inc(key)
}
time.Sleep(3 * time.Second)
fmt.Println(c.Value(key))
}