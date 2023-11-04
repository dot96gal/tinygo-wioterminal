package main

import (
	"machine"
	"time"
)

func main() {
	button1 := machine.BUTTON_1
	button1.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button1.SetInterrupt(machine.PinToggle, func(machine.Pin) {
		led.Set(button1.Get())
	})

	for {
		// heavy process
		time.Sleep(1 * time.Second)
	}
}
