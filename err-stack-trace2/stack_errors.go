package errstacktrace2

import (
	"errors"
	"fmt"
	"runtime"
)

type StackError struct {
	Stack []byte
	err   error
}

func NewStackError(err error) *StackError {
	var buf [16 * 1024]byte
	n := runtime.Stack(buf[:], false)
	return &StackError{
		err:   err,
		Stack: buf[:n],
	}
}

func (e *StackError) Error() string {
	return e.err.Error()
}

func (e *StackError) Unwrap() error {
	return e.err
}

func Wrap(err error, format string, args ...interface{}) error {
	if err != nil {
		err = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
		return err
	}
	return nil
}

func WrapStack(err error, format string, args ...interface{}) error {
	if err != nil {
		if se := (*StackError)(nil); !errors.As(err, &se) {
			err = NewStackError(err)
		}
		return Wrap(err, format, args...)
	}
	return nil
}
