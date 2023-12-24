package go_vmt

type BoneSetIndex int

const (
	RootAndCWrist BoneSetIndex = iota
	Thumb
	Index
	Middle
	Ring
	Pinky
)

type BoneIndex int

const (
	Root BoneIndex = iota
	Wrist
	Thumb0_ThumbProximal
	Thumb1_ThumbIntermediate
	Thumb2_ThumbDistal
	Thumb3_ThumbEnd
	IndexFinger0_IndexProximal
	IndexFinger1_IndexIntermediate
	IndexFinger2_IndexDistal
	IndexFinger3_IndexDistal2
	IndexFinger4_IndexEnd
	MiddleFinger0_MiddleProximal
	MiddleFinger1_MiddleIntermediate
	MiddleFinger2_MiddleDistal
	MiddleFinger3_MiddleDistal2
	MiddleFinger4_MiddleEnd
	RingFinger0_RingProximal
	RingFinger1_RingIntermediate
	RingFinger2_RingDistal
	RingFinger3_RingDistal2
	RingFinger4_RingEnd
	PinkyFinger0_PinkyProximal
	PinkyFinger1_PinkyIntermediate
	PinkyFinger2_PinkyDistal
	PinkyFinger3_PinkyDistal2
	PinkyFinger4_PinkyEnd
	Aux_Thumb_ThumbHelper
	Aux_Index_IndexHelper
	Aux_Middle_MiddleHelper
	Aux_Ring_RingHelper
	AuxVirtualSkeletonPinky_LittleHelper
)

type VirtualSkeleton struct {
	Index        int
	BoneSetIndex BoneSetIndex
	TimeOffset   TimeOffset
	Value        float64
	Position     Position
	Quaternion   Quaternion
	Euler        Euler
}

type SkeletonMessage string

const (
	SimpleSkeletonMessage     SkeletonMessage = "/VMT/Skeleton/Scalar"
	UnitySkeletonMessage      SkeletonMessage = "/VMT/Skeleton/Unity"
	UnitySkeletonEulerMessage SkeletonMessage = "/VMT/Skeleton/UEuler"
	OpenVRSkeletonMessage     SkeletonMessage = "/VMT/Skeleton/Driver"
	ApplySkeletonMessage      SkeletonMessage = "/VMT/Skeleton/Apply"
)
