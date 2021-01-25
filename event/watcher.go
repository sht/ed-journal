package event

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

type Watcher struct {
	watcher     *watcher.Watcher
	interval    time.Duration
	handlerFunc Handler
	journalRex  *regexp.Regexp
}

//func NewWatcher(h Handler, d time.Duration) (*Watcher, error) {
//	return &Watcher{
//		watcher:     watcher.New(),
//		interval:    d,
//		handlerFunc: h,
//	}, nil
//}

func (w *Watcher) Watch(path string) error {
	err := w.watcher.Add(path)
	if err != nil {
		return err
	}
	return nil
}

func (w *Watcher) Read(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	w.handlerFunc(b)
}

func (w *Watcher) Start() {
	fmt.Println("starting watcher")
	go func() {
		for {
			select {
			case event := <-w.watcher.Event:
				fmt.Println("READ")
				go w.Read(event.Path) // Print the event's info.
			case err := <-w.watcher.Error:
				fmt.Println(err)
				continue
			case <-w.watcher.Closed:
				return
			}
		}
	}()

	go func() {
		_ = w.watcher.Start(w.interval)
	}()
}

func (w *Watcher) Stop() {
	w.watcher.Close()
}
