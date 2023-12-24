package go_vmt

type DriverControl struct {
	ButtonIndex  int
	TriggerIndex int
	JointIndex   int
	TimeOffset   TimeOffset
	Serial       string
	Enable       int // 0: disable, 1: enable
}

type DriverControlMessage string

const (
	Reset                  DriverControlMessage = "/VMT/Reset"
	LoadSetting            DriverControlMessage = "/VMT/LoadSetting"
	SetRoomMatrix          DriverControlMessage = "/VMT/SetRoomMatrix"
	SetRoomMatrixTemporary DriverControlMessage = "/VMT/SetRoomMatrix/Temporary"
)
