package ev3

import (
	"testing"
)

func TestOn(t *testing.T) {
	ev3 := NewEventEmitter()
	ev3 = ev3.On("event", func() {})

	if _, ok := ev3.Events["event"]; !ok {
		t.Error("Failed to add listener to the ev3 emitter.")
	}
}

func TestRemove(t *testing.T) {
	ev3 := NewEventEmitter()
	ev3 = ev3.On("event", func() {}).Remove("event")

	if _, ok := ev3.Events["event"]; ok {
		t.Error("Failed to remove listener from emitter.")
	}
}

func TestReset(t *testing.T) {
	ev3 := NewEventEmitter()
	ev3.On("event1", func() {}).
		On("event2", func() {}).
		On("event3", func() {}).
		Reset()

	if len(ev3.Events) != 0 {
		t.Error("Failed to reset emitter.")
	}
}

func TestEmit(t *testing.T) {
	ev3 := NewEventEmitter()
	called := false
	cmdStart := func() { called = true }
	ev3.On("event1", cmdStart)
	ev3.Emit("event1")

	if !called {
		t.Errorf("Failed to emit event1")
	}
}
