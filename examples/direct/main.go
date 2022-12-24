//go:build macropad_rp2040

package main

import (
	"machine"
	"time"
)

var (
	old_AB      = 0b00000011
	states      = []int8{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}
	rotaryValue int
)

func interrupt(pin machine.Pin) {
	aHigh, bHigh := machine.ROT_A.Get(), machine.ROT_B.Get()
	old_AB <<= 2
	if aHigh {
		old_AB |= 1 << 1
	}
	if bHigh {
		old_AB |= 1
	}
	rotaryValue += int(states[old_AB&0x0f])
}

func main() {

	machine.ROT_A.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.ROT_A.SetInterrupt(machine.PinRising|machine.PinFalling, interrupt)

	machine.ROT_B.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.ROT_B.SetInterrupt(machine.PinRising|machine.PinFalling, interrupt)

	for oldValue := 0; ; {
		time.Sleep(100 * time.Microsecond) // doesn't work without this?
		if newValue := rotaryValue / 4; newValue != oldValue {
			println("value: ", newValue)
			oldValue = newValue
		}
	}

}
