package nimp

import (
    "fmt"
)

/*
Dropper is an interface that represents a resource that needs to be finalised
and its finalisation can fail.
*/
type Dropper interface {
    Drop() error
}

// Drop implements the Dropper interface for ErrorFn.
func (e ErrorFn) Drop() (err error) {
    return e()
}

/*
DropHandler is used to handle the Drop method of resources implementing
the Dropper interface.
*/
type DropHandler struct {
    tryUntil TryUntil
}

/*
NewDropHandler creates a DropHandler that wraps the tryUntil variable to
be used for closing resources.
*/
func NewDropHandler(tryUntil TryUntil) (c DropHandler) {
    return DropHandler {
        tryUntil: tryUntil,
    }
}

/*
Handle attempts to close the resource using the tryUntil variable passed
to the NewDropHandler function. If the error passed as the first argument
is nit nil, the error returned by exhaustion of tryUntil will be wrapped
in a DropError.
*/
func (d DropHandler) Handle(err error, errFn ErrorFn) (Err error) {
    exhaust := Exhaust(d.tryUntil, errFn)
    if err == nil {
        return exhaust
    }

    return DropError{ WhileDropping: exhaust, Cause: err }
}


/*
DropError contains two errors that occur when an error occurs while
using a Droppable resource (e.g. a closable file) and while trying to
drop the resource an error occurs.
*/
type DropError struct {
    // The error that occured when trying to close the resource.
    WhileDropping error
    // The error that occurred during the usage of a closable resource.
    Cause error
}

func (d DropError) Unwrap() (err error) {
    return d.Cause
}

func (d DropError) Error() (s string) {
    return fmt.Sprintf(
        "failed to drop resource (error: %s) while while handling error %s",
        d.WhileDropping.Error(), d.Cause.Error())
}
