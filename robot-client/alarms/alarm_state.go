package alarms

import (
	"container/list"
	"robot-client/logging"
)

type StateType int64

const (
	Ready StateType = iota
	Alarming
)

func (a StateType) String() string {
	switch a {
	case Ready:
		return "Ready"
	case Alarming:
		return "Alarming"
	}
	return "unknown"
}

var (
	currentAlarms = list.New()
	device_id     *string

	// FIXME hard coded location and temp
	// FIXME location and temp are pushed with all alarms
	// 		 need to add logging data from csv to generator

	location = logging.LocationClass{
		Latitude:  37.773972,
		Longitude: -122.431297,
	}
	temperature = 56.0
)

func Init(id *string) {
	device_id = id
}

func GetCurrentAlarm() AlarmName {
	if currentAlarms.Len() == 0 {
		return AlarmName_None
	}
	return currentAlarms.Front().Value.(*Alarm).name
}

func RaiseAlarm(alarm_name AlarmName) {
	entry := logging.LogEntry{
		DeviceID:    *device_id,
		Title:       "Raise Alarm",
		Entry:       alarm_name.String(),
		Location:    &location,
		Temperature: &temperature,
	}
	logging.Log(entry)
	alarm := NewAlarm(alarm_name)
	if currentAlarms.Len() == 0 {
		currentAlarms.PushBack(alarm)
		return
	}

	search := currentAlarms.Back()

	if !alarm.GreaterThan(search.Value.(*Alarm)) {
		currentAlarms.PushBack(alarm)
		return
	} else {
		for alarm.GreaterThan(search.Value.(*Alarm)) {
			search = search.Prev()
			if search == nil {
				currentAlarms.PushFront(alarm)
				return
			}
		}
		currentAlarms.InsertAfter(alarm, search)
	}

}

func ClearAlarm(alarm_name AlarmName) {

	search := currentAlarms.Front()
	for search != nil {
		if alarm_name == search.Value.(*Alarm).name {
			delete := search
			search = search.Next()
			currentAlarms.Remove(delete)
			entry := logging.LogEntry{
				DeviceID:    *device_id,
				Title:       "Clear Alarm",
				Entry:       alarm_name.String(),
				Location:    &location,
				Temperature: &temperature,
			}
			logging.Log(entry)
		} else {
			search = search.Next()
		}
	}

}
