package nimp

import (
    "testing"
)

type TestTryNTimes struct {
    try TryNTimes
    fn ErrorFn
    expected Status
}

func (t TestTryNTimes) TestWith() (err error) {
    Exhaust(t.try, t.fn)


    return
}
