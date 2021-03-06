package io

import (
    "github.com/bozso/gotoolbox/hash"
)

type ReaderRecord interface {
    ReaderCreator
    hash.Hashable
}

type WriterRecord interface {
    WriterCreator
    hash.Hashable
}

type ReaderDB interface {
    UseReader(hash.ID64, ReaderFn) error
    AddReader(ReaderRecord) hash.ID64
    SetReader(hash.ID64, ReaderRecord)
    CloseReader(hash.ID64)
}

type WriterDB interface {
    UseWriter(hash.ID64, WriterFn) error
    AddWriter(WriterRecord) hash.ID64
    SetWriter(hash.ID64, WriterRecord)
    CloseWriter(hash.ID64)
}

type IOService interface {
    ReaderDB
    WriterDB
    Close() error
}
