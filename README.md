# goframework-gorm-sqlite

包装`gorm-sqlite`对象，实现[`godb`](https://github.com/kordar/godb)接口。

## 安装
```go
go get github.com/kordar/goframework-gorm-sqlite@latest
```

## 使用

- 初始化 slog（可选，默认使用 slog.Default）

```go
import (
	"log/slog"
	"os"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	// 或文本输出
	// slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}
```

- 设置 GORM 日志等级（error | warn | info）

```go
import (
	goframeworkgormsqlite "github.com/kordar/goframework-gorm-sqlite"
)

logLevel := "info"  // error warn info
goframeworkgormsqlite.SetDbLogLevel(logLevel)
```

- 添加实例

```go
import (
	"log/slog"
	goframeworkgormsqlite "github.com/kordar/goframework-gorm-sqlite"
)

key := "default"
dsn := "file:test.db?cache=shared&_fk=1"
if err := goframeworkgormsqlite.AddSqliteInstance(key, dsn); err != nil {
	slog.Error("add sqlite instance failed", "err", err)
}
```

## 说明
- 已移除 gologger，GORM 日志通过标准库 slog 输出，需要 Go 1.21+
- 通过 SetDbLogLevel 控制 GORM SQL 日志级别；应用层可自由配置 slog 的 Handler/Level/Formatter

