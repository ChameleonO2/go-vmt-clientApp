package go_vmt

import "github.com/hypebeast/go-osc/osc"

type VMTClient struct {
	oscClient *osc.Client
	Trackers  []*VirtualTracker
}

func NewVMTClient(o *osc.Client) *VMTClient {
	return &VMTClient{
		oscClient: o,
		Trackers:  []*VirtualTracker{},
	}
}

func (c *VMTClient) AddTracker(t *VirtualTracker) {
	c.Trackers = append(c.Trackers, t)
}
func (c *VMTClient) SendAll() {
	for _, t := range c.Trackers {
		err := c.oscClient.Send(t.UnityRoomMessage())
		if err != nil {
			panic(err)
		}
	}
}
