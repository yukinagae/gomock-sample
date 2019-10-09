# WIP: gomock sample

Note: not working at the moment due to gomock `import cycle not allowed in test` error.

see: [mockgen source mode does not work with go1.11 modules]( https://github.com/golang/mock/issues/230#issuecomment-446743550)

## Usage

* Generate a mock from user.go

```bash
mockgen -source user.go -destination mock_user.go
```

* Run

```bash
go run main.go user.go
```

* Test

```bash
$ go test -self_package ./...
?   github.com/yukinagae/gomock-sample [no test files]
```
