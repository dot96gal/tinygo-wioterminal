package main

import (
	"machine"
)

func main() {
	button1 := machine.BUTTON_1
	button1.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Set(button1.Get())
	}
}
