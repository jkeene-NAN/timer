package main

import (
	"fmt"
	"github.com/timer"
	"io"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "test timer ", log.LstdFlags)
}

func main() {
	stream := os.Stdout

	//runStraightTimer(stream)
	runTempoChange(stream)

}

func runTempoChange(stream io.Writer) {
	type tempo struct {
		bpm   int
		ticks timer.TickType
	}

	tempos := []tempo{
		tempo{
			bpm:   60,
			ticks: 8,
		},
		tempo{
			bpm:   120,
			ticks: 4,
		},
		tempo{
			bpm:   180,
			ticks: 4,
		},
		tempo{
			bpm:   60,
			ticks: 8,
		},
		tempo{
			bpm:   30,
			ticks: 4,
		},
	}

	t := timer.CreateSystemClockTimer(timer.BeatsPerMinuteToNanoSecondsPerBeat(60))

	for _, temp := range tempos {
		t = t.SetNanoSecondsPerTick(timer.BeatsPerMinuteToNanoSecondsPerBeat(temp.bpm))
		io.WriteString(stream, fmt.Sprintf("tempo: %d\n", temp.bpm))
		ticksTotal := timer.TickType(0)
		ticksSince := timer.TickType(0)

		for ticksTotal < temp.ticks {
			ticksSince, ticksTotal, t = t.GetTicks()
			if ticksSince > 0 {
				io.WriteString(stream, fmt.Sprintf("ticksTotal: %d, ticksSince: %d\n", ticksSince, ticksTotal))
			}
		}
	}
}

func runStraightTimer(stream io.Writer) {
	runTimer(stream, 60, 8)
	runTimer(stream, 30, 8)
	runTimer(stream, 60, 8)
	runTimer(stream, 120, 8)
	runTimer(stream, 240, 8)
	runTimer(stream, 360, 8)
}

func runTimer(stream io.Writer, bpm int, ticks int) {
	logger.Printf("running timer at %d bpm for %d ticks\n", bpm, ticks)
	toStop := timer.TickType(ticks)
	totalTicks := timer.TickType(0)
	ticksSince := timer.TickType(0)
	t := timer.CreateSystemClockTimer(timer.BeatsPerMinuteToNanoSecondsPerBeat(bpm))

	for totalTicks < toStop {
		ticksSince, totalTicks, t = t.GetTicks()
		if ticksSince > 0 {
			io.WriteString(stream, fmt.Sprintf("total ticks: %d\n", totalTicks))
		}
	}
}
