package errstacktrace2_test

import (
	"fmt"
	"testing"

	errstacktrace2 "github.com/okatti-et14/go-examples/err-stack-trace2"
)

func TestWrapStack(t *testing.T) {
	fmt.Println(errstacktrace2.Fuga())
	t.Error(errstacktrace2.Fuga())
}
