package io

import (
    "io"

    "github.com/bozso/nimp"
)

// WriterCreator represents a variable that can be used to create a Writer.
type WriterCreator interface {
    CreateWriter() (io.WriteCloser, error)
}

// WriterFn represents a function that can use a Writer resource.
type WriterFn func(io.Writer) error

/*
UseWriter is a function for managing the creation usage and closing of a
readable resource.
*/
func UseWriter(d nimp.DropHandler, w WriterCreator, fn WriterFn) (err error) {
    rc, err := w.CreateWriter()
    if err != nil {
        return
    }

    return d.Handle(fn(rc), rc.Close)
}

