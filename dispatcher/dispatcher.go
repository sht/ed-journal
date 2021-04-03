package dispatcher

import (
	"github.com/sht/ed-journal/event"
)

var defaultDispatcher *Dispatcher

type Dispatcher struct {
	events map[string][]event.Handler
}

func New() *Dispatcher {
	return &Dispatcher{
		events: make(map[string][]event.Handler),
	}
}

func On(name string, h event.Handler) {
	if defaultDispatcher == nil {
		defaultDispatcher = New()
	}
	defaultDispatcher.On(name, h)
}

func Trigger(name string, b []byte) {
	if defaultDispatcher == nil {
		defaultDispatcher = New()
	}
	defaultDispatcher.Trigger(name, b)
}

func (d *Dispatcher) On(name string, h event.Handler) {
	_, ok := d.events[name]
	if !ok {
		d.events[name] = make([]event.Handler, 0, 1)
	}
	d.events[name] = append(d.events[name], h)
}

func (d *Dispatcher) Trigger(name string, b []byte) {
	handlers, ok := d.events[name]
	if !ok {
		//fmt.Printf("%s event is not registered\n", name)
		return
	}
	for _, h := range handlers {
		go h(b)
	}
}
