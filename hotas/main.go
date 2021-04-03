package main

import (
	"context"
	"fmt"
	"github.com/google/gousb"
	"github.com/sht/ed-journal/dispatcher"
	"github.com/sht/ed-journal/event"
	"github.com/sht/ed-journal/hotas/stick"
	"github.com/sht/ed-journal/hotas/throttle"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	ThrottleEvent             = "ThrottleEvent"
	StickEvent                = "StickEvent"
	ThrottleStateUpdatedEvent = "ThrottleStateUpdatedEvent"
	StickStateUpdatedEvent    = "StickStateUpdatedEvent"
	ThrottleConnectedEvent    = "ThrottleConnectedEvent"
	StickConnectedEvent       = "StickConnectedEvent"
)

var debug = true

func main() {
	log.SetOutput(ioutil.Discard)

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	dispatcher.On(ThrottleEvent, ThrottleEventHandler)
	dispatcher.On(StickEvent, StickEventHandler)

	if debug {
		dispatcher.On(ThrottleStateUpdatedEvent, func(b []byte) {
			//state := throttle.GetState()
			fmt.Println(ThrottleStateUpdatedEvent)
		})

		dispatcher.On(StickStateUpdatedEvent, func([]byte) {
			//state := stick.GetState()
			fmt.Println(StickStateUpdatedEvent)
		})
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)

	go setupDevice(ctx, &wg, DeviceOptions{
		Name:      "throttle",
		VendorID:  0x0738,
		ProductID: 0xa221,
		OnUpdate: func(b []byte) {
			dispatcher.Trigger(ThrottleStateUpdatedEvent, b)
		},
		OnConnect: func(b []byte) {
			fmt.Println(ThrottleConnectedEvent)
		},
	})

	go setupDevice(ctx, &wg, DeviceOptions{
		Name:      "stick",
		VendorID:  0x0738,
		ProductID: 0x2221,
		OnUpdate: func(b []byte) {
			dispatcher.Trigger(StickStateUpdatedEvent, b)
		},
		OnConnect: func(b []byte) {
			fmt.Println(StickConnectedEvent)
		},
	})

	<-stop
	fmt.Println()
	cancel()
	wg.Wait()
}

func ThrottleEventHandler(b []byte) {
	throttle.UpdateState(b)
	dispatcher.Trigger(ThrottleStateUpdatedEvent, nil)
}

func StickEventHandler(b []byte) {
	stick.UpdateState(b)
	dispatcher.Trigger(StickStateUpdatedEvent, nil)
}

type DeviceOptions struct {
	Name      string
	VendorID  gousb.ID
	ProductID gousb.ID
	OnUpdate  event.Handler
	OnConnect event.Handler
}

func setupDevice(c context.Context, wg *sync.WaitGroup, opt DeviceOptions) {
	defer wg.Done()
	ctx := gousb.NewContext()
	defer ctx.Close()
	ctx.Debug(0)

	dev, err := ctx.OpenDeviceWithVIDPID(opt.VendorID, opt.ProductID)
	if err != nil {
		log.Fatal(err)
	}
	if dev == nil {
		return
	}
	defer dev.Close()

	err = dev.SetAutoDetach(true)
	if err != nil {
		log.Fatal(err)
	}

	// Switch the configuration to #2.
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(2): %v", dev, err)
	}
	defer cfg.Close()

	// In the config #2, claim interface #3 with alt setting #0.
	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("%s.Interface(3, 0): %v", cfg, err)
	}
	defer intf.Close()

	// In this interface open endpoint #6 for reading.
	epIn, err := intf.InEndpoint(1)
	if err != nil {
		log.Fatalf("%s.InEndpoint(6): %v", intf, err)
	}

	stream, err := epIn.NewStream(epIn.Desc.MaxPacketSize, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	b := make([]byte, epIn.Desc.MaxPacketSize)

	opt.OnConnect(nil)

	go func() {
		for {
			readBytes, err := stream.Read(b)
			if err == io.ErrClosedPipe {
				return
			}
			if err != nil {
				fmt.Println("Read returned an error:", err)
				continue
			}
			if readBytes != 0 {
				opt.OnUpdate(b)
			}
		}
	}()

	<-c.Done()
}
