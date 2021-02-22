package io

import (
    "io"
    "fmt"

    "github.com/bozso/nimp"
)

/*
CloseHandler is used to handle the Close method of resources implementing
the Closer interface.
*/
type CloseHandler struct {
    tryUntil nimp.TryUntil
}

/*
NewCloseHandler creates a CloseHandler that wraps the tryUntil variable to
be used for closing resources.
*/
func NewCloseHandler(tryUntil nimp.TryUntil) (c CloseHandler) {
    return CloseHandler {
        tryUntil: tryUntil,
    }
}

/*
Handle attempts to close the resource using the tryUntil variable passed
to the NewCloseHandler function. If the error passed as the first argument
is nit nil, the error returned by exhaustion of tryUntil will be wrapped
in a CloseError.
*/
func (c CloseHandler) Handle(err error, closer io.Closer) (Err error) {
    exhaust := nimp.Exhaust(c.tryUntil, closer.Close)
    if err == nil {
        return exhaust
    }

    return CloseError{ WhileClosing: exhaust, Cause: err }
}


/*
CloseError contains two errors that occur when an error occurs while
using a Closable resource (e.g. a file) and while trying to close a file
an erro occurs.
*/
type CloseError struct {
    // The error that occured when trying to close the resource.
    WhileClosing error
    // The error that occurred during the usage of a closable resource.
    Cause error
}

func (c CloseError) Unwrap() (err error) {
    return c.Cause
}

func (c CloseError) Error() (s string) {
    return fmt.Sprintf(
        "failed to close resource (error: %s) while while handling error %s",
        c.WhileClosing.Error(), c.Cause.Error())
}

