package errstacktrace2

import "errors"

var ErrHoge error = errors.New("ErrHoge")

func Hoge() error {
	return WrapStack(ErrHoge, "Hoge")
}

func Fuga() error {
	return WrapStack(Hoge(), "Fuga")
}
