package io

import (
    "io"
)

type WriterCreator interface {
    CreateWriter() (io.WriteCloser, error)
}

type WriterFn func(io.Writer) error

type WriterUser interface {
    UseWriter(WriterCreator, WriterFn) error
}

func (c CloseHandler) UseWriter(w WriterCreator, fn WriterFn) (err error) {
    rc, err := w.CreateWriter()
    if err != nil {
        return
    }

    return c.Handle(fn(rc), rc)
}

