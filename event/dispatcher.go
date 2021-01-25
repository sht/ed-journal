package event

import (
	"fmt"
)

type Dispatcher struct {
	events map[string][]Handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		events: make(map[string][]Handler),
	}
}

func (d *Dispatcher) On(name string, h Handler) {
	_, ok := d.events[name]
	if !ok {
		d.events[name] = make([]Handler, 0, 1)
	}
	d.events[name] = append(d.events[name], h)
}

func (d *Dispatcher) Trigger(name string, b []byte) error {
	handlers, ok := d.events[name]
	if !ok {
		return fmt.Errorf("%s event is not registered", name)
	}
	for _, h := range handlers {
		go h(b)
	}
	return nil
}
