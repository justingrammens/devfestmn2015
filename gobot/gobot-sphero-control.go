package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"fmt"
	"strconv"
)

func main() {
	gbot := gobot.NewGobot()
	api.NewAPI(gbot).Start()

	spheros := map[string]string{
		"Sphero-BPO": "/dev/tty.Sphero-RWR-AMP-SPP",
	}

	for name, port := range spheros {
		spheroAdaptor := sphero.NewSpheroAdaptor("sphero", port)

		spheroDriver := sphero.NewSpheroDriver(spheroAdaptor, "sphero")

		work := func() {
			spheroDriver.SetRGB(uint8(255), uint8(0), uint8(0))
		}

		robot := gobot.NewRobot(name,
			[]gobot.Connection{spheroAdaptor},
			[]gobot.Device{spheroDriver},
			work,
		)
		robot.AddCommand("turn_blue", func(params map[string]interface{}) interface{} {
			spheroDriver.SetRGB(uint8(0), uint8(0), uint8(255))
			return nil
		})
		
		robot.AddCommand("turn_red", func(params map[string]interface{}) interface{} {
			spheroDriver.SetRGB(uint8(255), uint8(0), uint8(0))
			return nil
		})
		
		robot.AddCommand("turn_green", func(params map[string]interface{}) interface{} {
			spheroDriver.SetRGB(uint8(0), uint8(255), uint8(0))
			return nil
		})
		
		// use direction: XXX
		// use speed: XXX
		robot.AddCommand("drive", func(params map[string]interface{}) interface{} {
			fmt.Printf("Params: %+v\n", params)
			direction, _ := strconv.Atoi(params["direction"].(string))
			speed, _ := strconv.Atoi(params["speed"].(string))
			fmt.Printf("Direction: %d\n", direction)
			fmt.Printf("Speed: %d\n", speed)
			spheroDriver.Roll(uint8(speed), uint16(direction))
			return nil
		})

		gbot.AddRobot(robot)
	}

	gbot.Start()
}