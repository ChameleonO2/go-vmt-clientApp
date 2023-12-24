package main

import (
	"fmt"
	"math"
	"time"

	go_vmt "github.com/chameleono2/go-osc-tracker/go-vmt"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	client := osc.NewClient("localhost", 39570)
	tracker := go_vmt.NewTracker("VMT_0", 0, client)
	vmtClient := go_vmt.NewVMTClient(client)
	vmtClient.AddTracker(tracker)
	for {
		vmtClient.SendAll()
		tracker.Position.X = math.Sin(float64(time.Now().UnixMicro()) / 100000)
		tracker.Position.Y = math.Cos(float64(time.Now().UnixMicro()) / 100000)
		fmt.Printf("%s\n", tracker.UnityRoomMessage())
	}
}
