## Ev3 (WIP)

This is an easy to use event emitter in golang using go concurrency.

```go

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/thedhejavu/ev3"
)

func main() {
	ev3 := ev3.NewEventEmitter()

	ev3.On("lock", func(name string) {
		fmt.Println(name)
		time.Sleep(1000 * time.Millisecond)

	})

	ev3.On("unlock", func(name string) {
		fmt.Println(name)
	})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		ev3.Emit("lock", "hello world")
	}()
	go func() {
		defer wg.Done()
		ev3.Emit("unlock", "hello wwworld")
	}()
	go func() {
		defer wg.Done()
		ev3.Emit("lock", "hello wwwworld")
	}()

	wg.Wait()

	ev3.Remove("lock")
	ev3.Emit("lock", "this has been removed, so do nothing")
	ev3.Emit("unlock", "welome!!")
}

```

TODO

- Add Support for multiple events
- Add support for Custom URLs and File Path
