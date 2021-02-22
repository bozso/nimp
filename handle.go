package nimp

type ErrorFn func() error

type ErrorFnHandler interface {
    ErrorFnHandle(f ErrorFn) error
}
