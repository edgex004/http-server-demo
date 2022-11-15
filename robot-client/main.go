package main

import (
	"flag"
	"robot-client/alarms"
	"robot-client/logging"
	"time"
)

var (
	device_id  *string
	log_server *string
)

func init() {
	device_id = flag.String("device_id", "Robot-1", "identifier for robot")
	log_server = flag.String("log_server", "http://127.0.0.1:8000", "port number")
}

func main() {
	flag.Parse()
	alarms.Init(device_id)
	logging.Init(log_server)

	time.Sleep(5 * time.Second)

	alarms.RaiseAlarm(alarms.NetworkLost)
	time.Sleep(5 * time.Second)

	alarms.RaiseAlarm(alarms.MotorStall)
	time.Sleep(10 * time.Second)

	alarms.RaiseAlarm(alarms.PauseCommanded)
	time.Sleep(10 * time.Second)

	alarms.ClearAlarm(alarms.MotorStall)
	time.Sleep(10 * time.Second)

	alarms.ClearAlarm(alarms.NetworkLost)
	time.Sleep(10 * time.Second)

	alarms.ClearAlarm(alarms.PauseCommanded)
	time.Sleep(10 * time.Second)

}
