package go_vmt

import (
	"github.com/hypebeast/go-osc/osc"
)

type vmtMode int

const (
	Disable vmtMode = iota
	Tracker
	LeftController
	RightController
	TrackingReference
	IndexLeftController
	IndexRightController
	ViveTracker
)

type TrackerMessage string

const (
	UnityRoomMessage        TrackerMessage = "/VMT/Room/Unity"
	UnityRoomEulerMessage   TrackerMessage = "/VMT/Room/UEuler"
	OpenVRRoomMessage       TrackerMessage = "/VMT/Room/Driver"
	UnityRawMessage         TrackerMessage = "/VMT/Raw/Unity"
	UnityRawEulerMessage    TrackerMessage = "/VMT/Raw/UEuler"
	OpenVRRawMessage        TrackerMessage = "/VMT/Raw/Driver"
	UnityJointMessage       TrackerMessage = "/VMT/Joint/Unity"
	UnityJointEulerMessage  TrackerMessage = "/VMT/Joint/UEuler"
	OpenVRJointMessage      TrackerMessage = "/VMT/Joint/Driver"
	UnityFollowMessage      TrackerMessage = "/VMT/Follow/Unity"
	UnityFollowEulerMessage TrackerMessage = "/VMT/Follow/UEuler"
	OpenVRFollowMessage     TrackerMessage = "/VMT/Follow/Driver"
)

type TimeOffset float64

type VirtualTracker struct {
	client     *osc.Client
	Index      int32
	Enable     vmtMode
	TimeOffset TimeOffset
	Position   Position
	Quaternion Quaternion
	Euler      Euler
	Serial     string
}

func NewTracker(name string, index int32, c *osc.Client) *VirtualTracker {
	return &VirtualTracker{
		client:     c,
		Index:      index,
		Enable:     Tracker,
		TimeOffset: 0,
		Position:   Position{},
		Quaternion: Quaternion{},
		Euler:      Euler{},
		Serial:     name,
	}
}

func (v *VirtualTracker) UnityRoomMessage() *osc.Message {
	return osc.NewMessage(
		string(UnityRoomMessage),
		v.Index,
		int32(v.Enable),
		float32(v.TimeOffset),
		float32(v.Position.X),
		float32(v.Position.Y),
		float32(v.Position.Z),
		float32(v.Quaternion.QX),
		float32(v.Quaternion.QY),
		float32(v.Quaternion.QZ),
		float32(v.Quaternion.QW),
	)
}
