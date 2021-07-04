package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)
}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
    s.DurationSlept = duration
}

type Sleeper interface {
    Sleep()
}

type CountdownOperationsSpy struct {
    Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        sleeper.Sleep()
        fmt.Fprintln(out, i)
    }
    sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}

func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}