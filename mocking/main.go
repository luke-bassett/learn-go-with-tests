package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfugurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the definied Duration.
func (c *ConfugurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown prints a countdown from 3 to out with a delay between count provided by sleep.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfugurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
