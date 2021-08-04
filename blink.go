package main

import (
        "fmt"
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
        led := gpio.NewLedDriver(r, "7")

        gobot.Every(100*time.Millisecond, func() {
                led.Toggle()
                fmt.Println("Toggled LED")
        })
}
