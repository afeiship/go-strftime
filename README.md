# go-strftime
> Simple date format for golang.

## installation
```shell
go get -u github.com/afeiship/go-strftime
```

## usage
```shell
improt "github.com/afeiship/go-strftime"

fmt.Println(strftime.Format("datetime")) // 2021-08-01 10:10:10
fmt.Println(strftime.Format("date")) // 2021-08-01
fmt.Println(strftime.Format("time")) // 10:10:10
fmt.Println(strftime.Format("unix")) // 1627832610
```
