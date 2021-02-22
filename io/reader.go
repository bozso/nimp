package io

import (
    "io"

)

type ReaderCreator interface {
    CreateReader() (io.ReadCloser, error)
}


type ReaderFn func(io.Reader) error

type ReaderUser interface {
    UseReader(ReaderCreator, ReaderFn) error
}

func (c CloseHandler) UseReader(r ReaderCreator, fn ReaderFn) (err error) {
    rc, err := r.CreateReader()
    if err != nil {
        return
    }

    return c.Handle(fn(rc), rc)
}

