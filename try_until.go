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

type TryNTimes uint

func (t TryNTimes) CreateTry() (tu TryUntil) {
    return &TryNTimesImpl {
        current: 0,
        times: uint(t),
    }
}


type TryNTimesImpl struct {
    current uint
    times uint
}

func (t *TryNTimesImpl) Try(err error) (s Status) {
    if err == nil || t.current >= t.times {
        return Finished
    }

    t.current++
    return Continue
}

func Exhaust(t TryUntil, fn ErrorFn) (err error) {
    for {
        err = fn()
        if t.Try(err) == Finished {
            return
        }
    }
}
