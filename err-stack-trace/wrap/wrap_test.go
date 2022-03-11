package wrap

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	t.Run("エラーを1回Wrap", func(t *testing.T) {
		ErrHoge := errors.New("hoge")

		currentDir, _ := os.Getwd()
		fileLine := currentDir + "/hoge.go:4"
		function := "github.com/okatti-et14/go-examples/err-stack-trace/wrap.testCallWrap"

		trace := fmt.Sprintf("%s\n\t%s", function, fileLine)
		expect := fmt.Sprintf("%s\n%s", ErrHoge.Error(), trace)

		err := testCallWrap(ErrHoge)

		assert.Equal(t, expect, err.Error())
		assert.ErrorIs(t, err, ErrHoge)
	})

	t.Run("エラーを2回Wrap", func(t *testing.T) {
		ErrMonyo := errors.New("monyo")

		currentDir, _ := os.Getwd()
		fileLine := currentDir + "/hoge.go:4"
		function := "github.com/okatti-et14/go-examples/err-stack-trace/wrap.testCallWrap"

		fileLine2nd := currentDir + "/hoge.go:8"
		function2nd := "github.com/okatti-et14/go-examples/err-stack-trace/wrap.testCallWrap2nd"

		trace := fmt.Sprintf("%s\n\t%s", function, fileLine)
		trace2nd := fmt.Sprintf("%s\n\t%s", function2nd, fileLine2nd)
		expect := fmt.Sprintf("%s\n%s\n%s", ErrMonyo.Error(), trace, trace2nd)

		err := testCallWrap(ErrMonyo)
		err = testCallWrap2nd(err)

		assert.Equal(t, expect, err.Error())
		assert.ErrorIs(t, err, ErrMonyo)
	})

}
