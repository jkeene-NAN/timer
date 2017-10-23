package timer

import "time"

const NANO_SECONDS_PER_MINUTE TickType = TickType(60 * 1000000000)

func GetCurrentTimeInNanoSeconds() TickType {
	return TickType(time.Now().UnixNano())
}

func BeatsPerMinuteToNanoSecondsPerBeat(bpm int) TickType {
	aux := TickType(bpm)
	ret := NANO_SECONDS_PER_MINUTE / aux
	return ret
}
