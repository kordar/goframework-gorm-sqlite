# goframework-gorm-sqlite

包装`gorm-sqlite`对象，实现[`godb`](https://github.com/kordar/godb)接口。

## 安装
```go
go get github.com/kordar/goframework-gorm-sqlite v1.0.3
```

## 使用

- 设置日志等级

```go
logLevel := "info"  // error warn info
goframeworkgormsqlite.SetDbLogLevel(logLevel)
```

- 添加实例

```go
goframeworkgormsqlite.AddSqliteInstance(key, dsn)
```

