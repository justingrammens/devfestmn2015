package main

import (
        "fmt"
        "time"

        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
        gbot := gobot.NewGobot()

		direction := uint16(0)
		
        adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-RWR-AMP-SPP")
        spheroDriver := sphero.NewSpheroDriver(adaptor, "sphero")

        work := func() {

                gobot.Every(3*time.Second, func() {
                        spheroDriver.Roll(50, direction)
						fmt.Printf("Direction %+v\n", direction)
						
						if (direction >= 360) {
							direction = 0
						} else {
							direction += 90
						}
						
                        r := uint8(gobot.Rand(255))
                        g := uint8(gobot.Rand(255))
                        b := uint8(gobot.Rand(255))
                        spheroDriver.SetRGB(r, g, b)
						
						time.Sleep(1000 * time.Millisecond)
                })
        }

        robot := gobot.NewRobot("sphero",
                []gobot.Connection{adaptor},
                []gobot.Device{spheroDriver},
                work,
        )

        gbot.AddRobot(robot)

        gbot.Start()
}