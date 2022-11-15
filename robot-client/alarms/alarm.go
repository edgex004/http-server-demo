package alarms

func NewAlarm(a AlarmName) *Alarm {
	return &Alarm{name: a}
}

type Alarm struct {
	name AlarmName
}

func (a1 *Alarm) GreaterThan(a2 *Alarm) bool {
	priority1 := GetPriority(a1.name)
	priority2 := GetPriority(a2.name)

	return priority1 > priority2
}
