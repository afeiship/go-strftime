# go-strtime
> Simple date format for golang.

## installation
```shell
go get -u github.com/afeiship/go-strtime
```

## usage
```shell
improt "github.com/afeiship/go-strtime"

fmt.Println(strtime.Format("datetime")) // 2021-08-01 10:10:10
fmt.Println(strtime.Format("date")) // 2021-08-01
fmt.Println(strtime.Format("time")) // 10:10:10
fmt.Println(strtime.Format("unix")) // 1627832610
```
