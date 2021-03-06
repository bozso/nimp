package io

import (
    "io"

    "github.com/bozso/nimp"
)

// ReaderCreator represents a variable that can be used to create a Reader.
type ReaderCreator interface {
    CreateReader() (io.ReadCloser, error)
}

// ReaderFn represents a function that can use a Reader resource.
type ReaderFn func(io.Reader) error

/*
UseReader is a function for managing the creation usage and closing of a
readable resource.
*/
func UseReader(d nimp.DropHandler, r ReaderCreator, fn ReaderFn) (err error) {
    rc, err := r.CreateReader()
    if err != nil {
        return
    }

    return d.Handle(fn(rc), rc.Close)
}

