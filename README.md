# go-yoctog

too small log package wrapper

`log.Llongfile` フラグ立てると `calldepth` が 2 だから迂闊に `log` パッケージをラップ出来ないので誕生

# Usage

```bash
$ go get -u github.com/mobilusoss/go-yoctog
```

or

```bash
$ dep ensure -add github.com/mobilusoss/go-yoctog
```

# Example

```golang
log.SetFlags(log.Llongfile)
yoctog.Info("こんにちわーるど！")
```

```text
2018/12/04 14:51:58.718603 /path/to/main.go:52: [INFO] こんにちわーるど！
```

see example ( `example/main.go` ) file.
