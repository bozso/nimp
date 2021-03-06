package nimp

import (
    "os"
)

type InDir struct {
    current string
}

func (id InDir) ChangeBack() (err error) {
    return os.Chdir(id.current)
}

func ChangeDirectory(path string) (in InDir, err error) {
    if in.current, err = os.Getwd(); err != nil {
        return
    }

    err = os.Chdir(path)
    return
}

func InDirectory(d DropHandler, path string, fn ErrorFn) (err error) {
    id, err := ChangeDirectory(path)
    if err != nil {
        return
    }

    return d.Handle(fn(), id.ChangeBack)
}
