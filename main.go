package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/sht/ed-journal/event"
	"github.com/sht/ed-journal/events"
)

//func startWatcher() {
//	w, err := event.NewWatcher(func(b []byte) {
//		fmt.Println(string(b))
//	}, 10 * time.Millisecond)
//	if err != nil {
//		panic(err)
//	}
//	w.Start()
//	err = w.Watch("journal.log")
//	if err != nil {
//		panic(err)
//	}
//}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	d := event.NewDispatcher()

	// add event listeners
	events.AddListeners(d)

	b, err := ioutil.ReadFile("journal.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := bytes.Split(b, []byte("\n"))
	for _, b = range bs {
		if len(b) > 0 {
			go parseEvent(d, b)
		}
	}

	<-quit
}

func parseEvent(d *event.Dispatcher, b []byte) {
	var e event.Event
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	// trigger the event
	err = d.Trigger(e.Event, b)
	if err != nil {
		//fmt.Println(err)
		return
	}
}
