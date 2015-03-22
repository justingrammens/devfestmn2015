package main

import (
  "fmt"

  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot/platforms/sphero"
)

func main() {
  gbot := gobot.NewGobot()

  adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/rfcomm0")
  spheroDriver := sphero.NewSpheroDriver(adaptor, "sphero")

  work := func() {
    gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
      fmt.Println("Collision Detected!")
      r := uint8(gobot.Rand(255))
      g := uint8(gobot.Rand(255))
      b := uint8(gobot.Rand(255))
      spheroDriver.SetRGB(r, g, b)
    })
   }
   
   gobot.Every(3*time.Second, func() {
        spheroDriver.Roll(30, uint16(gobot.Rand(360)))
   		fmt.Printf("Roll it random direction\n")
	})
   
   
	
  robot := gobot.NewRobot("sphero",
    []gobot.Connection{adaptor},
    []gobot.Device{spheroDriver},
    work,
  )

  gbot.AddRobot(robot)

  gbot.Start()
}