# err-stack-trace
Simple Go error stack trace

run
```
err-stack-trace $ go run .
2022/03/12 01:11:43 hogehoge
main.stack1
        /{dir}/go-examples/err-stack-trace/main.go:17
main.stack2
        /{dir}/go-examples/err-stack-trace/main.go:23
2022/03/12 01:11:43 ErrHoge
github.com/okatti-et14/go-examples/err-stack-trace/hoge.Hoge
        /{dir}/go-examples/err-stack-trace/hoge/hoge.go:12
```
