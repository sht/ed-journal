package main

//func driver() {
//	stop := make(chan os.Signal, 2)
//	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//
//	dev, err := uinput.CreateTouchPad("/dev/uinput", []byte("Mad Catz Saitek Pro Flight X-56 Rhino Throttle"), [8]uinput.Axis{}, [36]uinput.Button{})
//	defer func() {
//		err = dev.Close()
//		if err != nil {
//			fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
//			os.Exit(1)
//		}
//	}()
//	if err != nil {
//		fmt.Printf("Failed to create the virtual mouse. Last error was: %s\n", err)
//		os.Exit(1)
//	}
//
//	var i int32 = 0
//	for {
//		i = (i + 1) % 1024
//		select {
//		case <-stop:
//			return
//		default:
//			time.Sleep(30 * time.Millisecond)
//			fmt.Println("Moving to ", i)
//			err = dev.SetAxis(uint16(throttle.AxisL), i)
//			if err != nil {
//				fmt.Printf("Failed to move mouse left. Last error was: %s\n", err)
//				os.Exit(1)
//			}
//		}
//	}
//}
