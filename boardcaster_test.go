package area_framework

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewBoardCaster(t *testing.T) {
	caster := NewBoardCaster()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		innerI := i
		go func() {
			receiver := caster.Listens()
			count := 0
			for {
				_ = receiver.Read()
				count++
				fmt.Printf("index:%d count:%d \n", innerI, count)
				if count == 10 {
					wg.Done()
				}
			}
		}()
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		caster.Write(&Message{
			msgType: Start,
		})
	}
	wg.Wait()
}
