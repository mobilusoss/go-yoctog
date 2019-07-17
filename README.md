# go-yoctog
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmobilusoss%2Fgo-yoctog.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmobilusoss%2Fgo-yoctog?ref=badge_shield)
[![Go Report](https://goreportcard.com/badge/github.com/mobilusoss/go-yoctog)](https://goreportcard.com/report/github.com/mobilusoss/go-yoctog)
[![codebeat badge](https://codebeat.co/badges/3e79f110-ec9b-4f5a-899a-1b5726f0a3b9)](https://codebeat.co/projects/github-com-mobilusoss-go-yoctog-master)
[![codecov](https://codecov.io/gh/mobilusoss/go-yoctog/branch/master/graph/badge.svg)](https://codecov.io/gh/mobilusoss/go-yoctog)
[![Documentation](https://godoc.org/github.com/mobilusoss/go-yoctog?status.svg)](https://godoc.org/github.com/mobilusoss/go-yoctog)

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
// Default log level are debug
//yoctog.SetLogLevel(yoctog.LevelDebug)
yoctog.Info("こんにちわーるど！")
```

```text
2018/12/04 14:51:58.718603 /path/to/main.go:52: [INFO] こんにちわーるど！
```

see example ( `example/main.go` ) file.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmobilusoss%2Fgo-yoctog.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmobilusoss%2Fgo-yoctog?ref=badge_large)