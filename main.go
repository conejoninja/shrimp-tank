package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/onewire"

	"tinygo.org/x/drivers/ds18b20"
)

var (
	temperature     int32
	sensors_ds18b20 ds18b20.Device
)

const (
	ledPin4 = machine.Pin(machine.D4)
	ledPin5 = machine.Pin(machine.D5)
	ledPin6 = machine.Pin(machine.D6)
)

func main() {
	ledPin4.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledPin5.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledPin6.Configure(machine.PinConfig{Mode: machine.PinOutput})
	temperature = 25000
	ow := onewire.New(machine.D2)
	romIDs, err := ow.Search(onewire.SEARCH_ROM)
	if err != nil {
		println(err)
	}
	sensors_ds18b20 := ds18b20.New(ow)
	for true {
		for _, romid := range romIDs {
			sensors_ds18b20.RequestTemperature(romid)
		}
		time.Sleep(1024 * time.Millisecond)
		temperature, _ = func() (tmp int32, err error) {
			for _, romid := range romIDs {
				tmp, err = sensors_ds18b20.ReadTemperature(romid)
				break
			}
			return
		}()

		ledPin4.Low()
		ledPin5.Low()
		ledPin6.Low()
		if temperature > 30000 {
			ledPin4.High()
		} else if temperature < 20000 {
			ledPin5.High()
		} else {
			ledPin6.High()
		}
		time.Sleep(1 * time.Second)
	}
}
