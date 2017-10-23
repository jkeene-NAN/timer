package timer_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/timer"
	"testing"
)

func TestCreateSystemClockTimer(t *testing.T) {
	Convey("When testing CreateSystemClockTimer", t, func() {
		Convey("It should populate the timer with the correct values", func() {
			bpm := 60
			nanoSecondsPerTick := timer.BeatsPerMinuteToNanoSecondsPerBeat(bpm)
			t := timer.CreateSystemClockTimer(nanoSecondsPerTick)

			Convey("The nano seconds per tick should equal the expected", func() {
				So(t.GetNanoSecondsPerTick(), ShouldEqual, nanoSecondsPerTick)
			})
		})
	})
}
