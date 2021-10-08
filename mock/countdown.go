package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const CountdownStart = 3

type Sleeper interface {
	sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		s.sleep()
		fmt.Fprintln(w, i)
	}
	s.sleep()
	fmt.Fprint(w, finalWord)

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
