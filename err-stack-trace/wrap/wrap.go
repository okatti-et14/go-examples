package wrap

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	// Get one stack trace
	pc := make([]uintptr, 1)
	// 0 identifying the frame for Callers itself and 1 identifying the caller of Callers.
	// 2 identifying the caller of Wrap.
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	// trace
	trace := fmt.Sprintf("%s\n\t%s:%d", frame.Function, frame.File, frame.Line)
	// stack trace error
	return fmt.Errorf("%w\n%s", err, trace)
}
