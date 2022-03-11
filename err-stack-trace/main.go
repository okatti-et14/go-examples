package main

import (
	"errors"
	"log"

	"github.com/okatti-et14/go-examples/err-stack-trace/hoge"
	"github.com/okatti-et14/go-examples/err-stack-trace/wrap"
)

func main() {
	log.Println(stack2())
	log.Println(hoge.Hoge())
}

func stack1() error {
	return wrap.Wrap(errors.New("hogehoge"))
}

func stack2() error {
	err := stack1()
	if err != nil {
		return wrap.Wrap(err)
	}
	return nil
}
