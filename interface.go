package timer

//TickType is the type returned from time.Now().Unixnano()
type TickType int64

//Timer is the interface for Timer objects
type Timer interface {

	//Set the nano seconds per tick.  Returns a new Timer
	SetNanoSecondsPerTick(nanoSeconds TickType) Timer

	//GetNanoSecondsPerTick returns the configured Nano seconds per tick
	GetNanoSecondsPerTick() TickType

	//GetTicks returns the number of ticks since this function was last called, the number of NanoSeconds
	//that have elapses since this function was last called, and a new Timer to be used in the next iteration
	GetTicks() (TickType, TickType, Timer)
}
