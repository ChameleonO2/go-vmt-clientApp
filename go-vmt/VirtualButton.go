package go_vmt

import "fmt"

type VirtualButton struct {
	Index         int
	ButtonIndex   int
	TriggerIndex  int
	JoyStickIndex int
	TimeOffset    TimeOffset
}

type ButtonMessage string

const (
	ButtonInputMessage        ButtonMessage = "/VMT/Input/Button"
	ButtonTouchInputMessage   ButtonMessage = "/VMT/Input/Button/Touch"
	TriggerInputMessage       ButtonMessage = "/VMT/Input/Trigger"
	TriggerTouchInputMessage  ButtonMessage = "/VMT/Input/Trigger/Touch"
	TriggerClickInputMessage  ButtonMessage = "/VMT/Input/Trigger/Click"
	JoyStickInputMessage      ButtonMessage = "/VMT/Input/JoyStick"
	JoyStickTouchInputMessage ButtonMessage = "/VMT/Input/JoyStick/Touch"
	JoyStickClickInputMessage ButtonMessage = "/VMT/Input/JoyStick/Click"
)

type PressValue int

const (
	ReleaseValueInt PressValue = 0
	PressValueInt   PressValue = 1
)

type PressMap map[PressBool]PressValue

type PressBool bool

const (
	Release PressBool = false
	Press   PressBool = true
)

var PressMapTable = PressMap{
	Release: ReleaseValueInt,
	Press:   PressValueInt,
}

func (v *VirtualButton) ButtonInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", ButtonInputMessage, v.Index, v.ButtonIndex, v.TimeOffset, PressMapTable[b])
}

func (v *VirtualButton) ButtonTouchInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", ButtonTouchInputMessage, v.Index, v.ButtonIndex, v.TimeOffset, PressMapTable[b])
}

type TriggerValue float64

func (v *VirtualButton) TriggerInputMessage(t TriggerValue) string {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return fmt.Sprintf("%s %d, %d , %f, %f", TriggerInputMessage, v.Index, v.TriggerIndex, v.TimeOffset, t)
}
func (v *VirtualButton) TriggerTouchInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", TriggerTouchInputMessage, v.Index, v.TriggerIndex, v.TimeOffset, PressMapTable[b])
}

func (v *VirtualButton) TriggerClickInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", TriggerClickInputMessage, v.Index, v.TriggerIndex, v.TimeOffset, PressMapTable[b])
}

type JoyStickValue struct {
	X float64
	Y float64
}

func (v *VirtualButton) JoyStickInputMessage(j JoyStickValue) string {
	return fmt.Sprintf("%s %d, %d , %f, %f, %f", JoyStickInputMessage, v.Index, v.JoyStickIndex, v.TimeOffset, j.X, j.Y)
}

func (v *VirtualButton) JoyStickTouchInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", JoyStickTouchInputMessage, v.Index, v.JoyStickIndex, v.TimeOffset, PressMapTable[b])
}

func (v *VirtualButton) JoyStickClickInputMessage(b PressBool) string {
	return fmt.Sprintf("%s %d, %d , %f, %d", JoyStickClickInputMessage, v.Index, v.JoyStickIndex, v.TimeOffset, PressMapTable[b])
}
