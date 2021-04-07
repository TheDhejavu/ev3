package ev3

import (
	"errors"
	"reflect"
	"sync"
)

type EventEmitter struct {
	Events map[string]interface{}
	*sync.Mutex
}

func NewEventEmitter() *EventEmitter {
	eventInstance := &EventEmitter{
		Mutex:  new(sync.Mutex),
		Events: make(map[string]interface{}),
	}
	return eventInstance
}

func (e *EventEmitter) On(name string, listener interface{}) *EventEmitter {
	e.Lock()
	defer e.Unlock()

	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		panic(errors.New("Listener must be a func."))
	}
	e.Events[name] = listener

	return e
}

func (e *EventEmitter) Remove(name string) *EventEmitter {
	e.Lock()
	defer e.Unlock()

	delete(e.Events, name)
	return e
}

func (e *EventEmitter) Reset() *EventEmitter {
	e.Lock()
	defer e.Unlock()

	e.Events = make(map[string]interface{})
	return e
}

func (e *EventEmitter) Emit(name string, args ...interface{}) *EventEmitter {
	listener := e.Events[name]

	e.Lock()

	if _, ok := e.Events[name]; !ok {
		e.Unlock()
		return e
	}
	
	defer e.Unlock()

	var wg sync.WaitGroup

	wg.Add(1)

	go func(listener interface{}) {
		defer wg.Done()

		fn := reflect.ValueOf(listener)
		rargs := make([]reflect.Value, len(args))
		for i, a := range args {
			rargs[i] = reflect.ValueOf(a)
		}
		fn.Call(rargs)

	}(listener)

	wg.Wait()
	return e
}
