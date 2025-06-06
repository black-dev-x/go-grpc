Generate proto files with
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

Testing it with evans

You can install evans with the terminal, but there are other approaches at their github
```
go install github.com/ktr0731/evans@latest
```

```
evans -r repl
```
