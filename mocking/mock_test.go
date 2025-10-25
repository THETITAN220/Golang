package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("print from 3 to 1", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
	t.Run("duration of sleep", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v, got %v", sleepTime, spyTime.durationSlept)
		}

	})
}
