package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type philosopher struct {
	ID int
}

func (p *philosopher) eating(mutex []sync.Mutex, i int) {
	mutex[i%5].Lock()
	defer mutex[i%5].Unlock()
	mutex[(i+1)%5].Lock()
	defer mutex[(i+1)%5].Unlock()
	fmt.Println(strconv.Itoa(p.ID) + "正在干饭...")
	time.Sleep(time.Second)
	fmt.Println(strconv.Itoa(p.ID) + "干饭完毕")
}

func main() {
	mutex := make([]sync.Mutex, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			pher := &philosopher{
				ID: i,
			}
			pher.eating(mutex, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
