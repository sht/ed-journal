package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/gousb"
	"github.com/rdnt/uinput"
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
	ThrottleEvent             = "Throttle"
	StickEvent                = "Stick"
	ThrottleStateUpdatedEvent = "ThrottleStateUpdated"
	StickStateUpdatedEvent    = "StickStateUpdated"
	ThrottleConnectedEvent    = "ThrottleConnected"
	StickConnectedEvent       = "StickConnected"
)

var debug = true

func main() {
	log.SetOutput(ioutil.Discard)

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	dispatcher.On(ThrottleEvent, ThrottleEventHandler)
	dispatcher.On(StickEvent, StickEventHandler)

	throttleDev, err := uinput.CreateJoystick(
		"/dev/uinput",
		[]byte("Mad Catz Saitek Pro Flight X-56 Rhino Throttle (emulated)"),
		throttle.AxesConfig,
		throttle.ButtonsConfig,
	)
	defer func() {
		err = throttleDev.Close()
		if err != nil {
			fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
			os.Exit(1)
		}
	}()
	if err != nil {
		fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
		os.Exit(1)
	}

	stickDev, err := uinput.CreateJoystick(
		"/dev/uinput",
		[]byte("Mad Catz Saitek Pro Flight X-56 Rhino Stick (emulated)"),
		stick.AxesConfig,
		stick.ButtonsConfig,
	)
	defer func() {
		err = stickDev.Close()
		if err != nil {
			fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
			os.Exit(1)
		}
	}()
	if err != nil {
		fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
		os.Exit(1)
	}

	dispatcher.On(ThrottleStateUpdatedEvent, func(_ []byte) {
		axes, buttons := throttle.GetState().Map()

		for id, axis := range axes {
			err = throttleDev.SetAxis(id, int32(axis))
			if err != nil {
				fmt.Printf("Failed to move mouse left. Last error was: %s\n", err)
				os.Exit(1)
			}
		}

		for id, on := range buttons {
			err = throttleDev.SetButton(id, bool(on))
			if err != nil {
				fmt.Printf("Failed to move mouse left. Last error was: %s\n", err)
				os.Exit(1)
			}
		}
	})

	dispatcher.On(StickStateUpdatedEvent, func(_ []byte) {
		axes, buttons := stick.GetState().Map()

		for id, axis := range axes {
			err = stickDev.SetAxis(id, int32(axis))
			if err != nil {
				fmt.Printf("Failed to move mouse left. Last error was: %s\n", err)
				os.Exit(1)
			}
		}

		for id, on := range buttons {
			err = stickDev.SetButton(id, bool(on))
			if err != nil {
				fmt.Printf("Failed to move mouse left. Last error was: %s\n", err)
				os.Exit(1)
			}
		}
	})

	if debug {
		dispatcher.On(ThrottleStateUpdatedEvent, func(b []byte) {
			//state := stick.GetState()
			//fmt.Println(ThrottleStateUpdatedEvent, state)
		})

		dispatcher.On(StickStateUpdatedEvent, func(b []byte) {
			for _, b := range b {
				fmt.Printf("%08b\n", b)
			}
			state := stick.GetState()
			b, _ = json.MarshalIndent(state, "", "  ")
			fmt.Println(string(b))
			//fmt.Println(StickStateUpdatedEvent, state)
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
			throttle.UpdateState(b)
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
			stick.UpdateState(b)
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
