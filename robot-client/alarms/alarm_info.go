//
// This file was generated by ./scripts/generate_code.py
// Do not modify directly. Modify alarm_info.csv and re-generate.
//

package alarms


type AlarmName int64

const (
	AlarmName_None AlarmName = iota
	MotorStall
	FirmwareLoadFail
	UpdateFailed
	PauseCommanded
	NetworkLost
)

func (a AlarmName) String() string {
	switch a {
	case AlarmName_None:
		return "None"
	case MotorStall:
		return "MotorStall"
	case FirmwareLoadFail:
		return "FirmwareLoadFail"
	case UpdateFailed:
		return "UpdateFailed"
	case PauseCommanded:
		return "PauseCommanded"
	case NetworkLost:
		return "NetworkLost"
	}
	return "unknown"
}

type Priority int64

const (
	Priority_None Priority = iota
	Low
	Medium
	High
)

func (a Priority) String() string {
	switch a {
	case Priority_None:
		return "None"
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	}
	return "unknown"
}

type LedIndicator int64

const (
	LedIndicator_None LedIndicator = iota
	RedFlash
	YellowSolid
	YellowFlash
	OrangeFlash
)

func (a LedIndicator) String() string {
	switch a {
	case LedIndicator_None:
		return "None"
	case RedFlash:
		return "RedFlash"
	case YellowSolid:
		return "YellowSolid"
	case YellowFlash:
		return "YellowFlash"
	case OrangeFlash:
		return "OrangeFlash"
	}
	return "unknown"
}

type Sound int64

const (
	Sound_None Sound = iota
	Beep
	Blare
)

func (a Sound) String() string {
	switch a {
	case Sound_None:
		return "None"
	case Beep:
		return "Beep"
	case Blare:
		return "Blare"
	}
	return "unknown"
}

func GetLedIndicator(a AlarmName) LedIndicator {
	switch a {
	case MotorStall:
		return RedFlash
	case FirmwareLoadFail:
		return RedFlash
	case UpdateFailed:
		return YellowSolid
	case PauseCommanded:
		return YellowFlash
	case NetworkLost:
		return OrangeFlash
	}
	return LedIndicator_None
}

func GetSound(a AlarmName) Sound {
	switch a {
	case MotorStall:
		return Blare
	case FirmwareLoadFail:
		return Blare
	case UpdateFailed:
		return Beep
	case PauseCommanded:
		return Beep
	case NetworkLost:
		return Beep
	}
	return Sound_None
}

func GetPriority(a AlarmName) Priority {
	switch a {
	case MotorStall:
		return High
	case FirmwareLoadFail:
		return High
	case UpdateFailed:
		return Low
	case PauseCommanded:
		return Low
	case NetworkLost:
		return Medium
	}
	return Priority_None
}