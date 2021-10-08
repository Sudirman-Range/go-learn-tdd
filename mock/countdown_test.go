package main

import (
	"bytes"
	"testing"
)

//Sleeperインターフェースのmock
type SleepCount struct {
	Count int
}

func (c *SleepCount) sleep() {
	c.Count++
}

func TestCountdown(t *testing.T) {
	sleepCount := &SleepCount{}
	buffer := &bytes.Buffer{}
	Countdown(buffer, sleepCount)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if sleepCount.Count != 4 {
		t.Errorf("not enough calls to sleeper, got %q want 4", sleepCount.Count)
	}
}
