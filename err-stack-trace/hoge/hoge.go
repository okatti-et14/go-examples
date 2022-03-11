package hoge

import (
	"errors"

	"github.com/okatti-et14/go-examples/err-stack-trace/wrap"
)

var ErrHoge error = errors.New("ErrHoge")

func Hoge() error {
	return wrap.Wrap(ErrHoge)
}
