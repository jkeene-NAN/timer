package timer

type systemClockTimer struct {
	nanoSecondsPerTick TickType
	inception          TickType
	ticks              TickType
}

//Set the nano seconds per tick.  Returns a new Timer with the old nanoSeconds value and a new nano seconds per tick
func (t systemClockTimer) SetNanoSecondsPerTick(nanoSeconds TickType) Timer {

	newNanoSecondsPerTick := nanoSeconds
	now := GetCurrentTimeInNanoSeconds()
	lastTickTime := (t.ticks * t.nanoSecondsPerTick) + t.inception
	newTicksSinceLastTickTime := (now - lastTickTime) / newNanoSecondsPerTick

	ret := systemClockTimer{
		inception:          now,
		nanoSecondsPerTick: newNanoSecondsPerTick,
		ticks:              newTicksSinceLastTickTime,
	}

	return ret
}

//GetNanoSecondsPerTick returns the configured Nano seconds per tick
func (t systemClockTimer) GetNanoSecondsPerTick() TickType {
	return t.nanoSecondsPerTick
}

//GetTicks returns the number of ticks since this function was last called, the number of NanoSeconds
//that have elapses since this function was last called, and a new Timer to be used in the next iteration
func (t systemClockTimer) GetTicks() (TickType, TickType, Timer) {
	now := GetCurrentTimeInNanoSeconds()
	elapsedTime := now - t.inception
	ticks := elapsedTime / t.nanoSecondsPerTick
	diff := ticks - t.ticks

	t.ticks = t.ticks + diff
	return diff, ticks, t
}

func CreateSystemClockTimer(nanoSecondsPerTick TickType) Timer {
	timer := systemClockTimer{
		inception:          GetCurrentTimeInNanoSeconds(),
		nanoSecondsPerTick: nanoSecondsPerTick,
		ticks:              0,
	}

	return timer
}
