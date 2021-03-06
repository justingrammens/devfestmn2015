package main

import (
        "fmt"
        "time"

        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
        gbot := gobot.NewGobot()

        adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-RWR-AMP-SPP")
        spheroDriver := sphero.NewSpheroDriver(adaptor, "sphero")

        work := func() {
                gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
                        fmt.Printf("Collision Detected! %+v\n", data)
                })

                gobot.Every(3*time.Second, func() {
                        spheroDriver.Roll(30, uint16(gobot.Rand(360)))
						fmt.Printf("Roll it random direction\n")
                })

                gobot.Every(1*time.Second, func() {
                        r := uint8(gobot.Rand(255))
                        g := uint8(gobot.Rand(255))
                        b := uint8(gobot.Rand(255))
                        spheroDriver.SetRGB(r, g, b)
						fmt.Printf("Set a random color\n")
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