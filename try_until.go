package nimp

import (

)

type Status bool

const (
    Continue Status = false
    Finished Status = true
)

type TryCreator interface {
    CreateTry() TryUntil
}

type TryUntil interface {
    Try(error) Status
}

type TryNTimes struct {
    current uint
    times uint
}

func (t *TryNTimes) Try(err error) (s Status) {
    if err == nil || t.current >= t.times {
        return Finished
    }

    t.current++
    return Continue
}

func Exhaust(t TryUntil, fn ErrorFn) (err error) {
    err = fn()
    for t.Try(err) != Finished {
        err = fn()
    }

    return
}
