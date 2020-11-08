# [WIP]: cahsper-cli

CLI tool for [Cahsper](https://github.com/YoshinoriN/cahsper).

# Requirements

* go 1.15

# Commands

|command||
|---|---|
|`help`, `-h`, `--help`|Help about any command.|
|`version`|Show version number.|
|`init`|Initialize cahsper-cli.|
|`credential list`|Show cahsper credential variables.|
|`credential set`|Set cahsper credential variables.|
|`config list`|Show cahsper configure variables.|
|`config set`|Set cahsper configure variables.|

# Build

```sh
$ go build
```

# Test

```sh
$ go test ./...

// with coverage
$ go test ./... -v -cover
```