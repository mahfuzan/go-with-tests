package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows us to put delays
type Sleeper interface {
	Sleep()
}

// Sleep with configurable duration
type ConfigurableSleeper struct {
	duration time.Duration       // to configure the duration of sleep
	sleep    func(time.Duration) // way to pass in a sleep function
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
