package nimp

import (
    "errors"
    "testing"
    "reflect"
)

type produceNErrors struct {
    ntimes uint
    current uint
}

var produced = errors.New("produced error")

func newProducer(ntimes uint) (p *produceNErrors) {
    return &produceNErrors {
        ntimes: ntimes,
        current: 0,
    }
}

func (p *produceNErrors) Call() (err error) {
    err = nil
    if p.current >= p.ntimes {
        return
    }

    p.current++
    return produced
}


type TryNTimesTest struct {
    createTry TryNTimes
    fn ErrorFn
    expected error
}

func (t TryNTimesTest) Test(test *testing.T) {
    err := Exhaust(t.createTry.CreateTry(), t.fn)

    if exp := t.expected; !reflect.DeepEqual(err, exp) {
        test.Fatalf("expected error '%s', got '%s'", exp, err)
    }
}

func TestTryNTimes(t *testing.T) {
    try := []TryNTimesTest{
        {TryNTimes(5), newProducer(5).Call, nil},
        {TryNTimes(5), newProducer(6).Call, produced},
    }

    for _, tryer := range try {
        tryer.Test(t)
    }
}
