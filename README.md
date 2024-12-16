# GoNANP
[![Go Report Card](https://goreportcard.com/badge/github.com/DLzer/gonanp)](https://goreportcard.com/report/github.com/DLzer/gonanp)
[![Test](https://github.com/DLzer/gonanp/actions/workflows/test.yml/badge.svg)](https://github.com/DLzer/gonanp/actions/workflows/test.yml)

Generate North American Numbering Plan compatible phone numbers ex `7725059292`

## Install
```bash
go get github.com/DLzer/gonanp
```

## Generator usage
```go
import "github.com/DLzer/gonanp

number := gonanp.GenerateNanp()
```

## Validator usage
```go
import "github.com/DLzer/gonanp

number := 7725059292
isValid := gonanp.ValidateNanp(number)
```