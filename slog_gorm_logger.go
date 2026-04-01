package goframework_gorm_sqlite

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type slogGormLogger struct {
	base           *slog.Logger
	level          logger.LogLevel
	slowThreshold  time.Duration
	ignoreNotFound bool
}

func newSlogGormLogger(level string) logger.Interface {
	lv := logger.Info
	if level == "error" {
		lv = logger.Error
	}
	if level == "warn" {
		lv = logger.Warn
	}
	return &slogGormLogger{
		base:           slog.Default(),
		level:          lv,
		slowThreshold:  200 * time.Millisecond,
		ignoreNotFound: true,
	}
}

func (l *slogGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	cp := *l
	cp.level = level
	return &cp
}

func (l *slogGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Info {
		return
	}
	l.base.InfoContext(ctx, fmt.Sprintf(msg, data...))
}

func (l *slogGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Warn {
		return
	}
	l.base.WarnContext(ctx, fmt.Sprintf(msg, data...))
}

func (l *slogGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level < logger.Error {
		return
	}
	l.base.ErrorContext(ctx, fmt.Sprintf(msg, data...))
}

func (l *slogGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level == logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	args := []any{"elapsed", elapsed, "sql", sql}
	if rows != -1 {
		args = append(args, "rows", rows)
	}
	if err != nil && !(l.ignoreNotFound && errors.Is(err, gorm.ErrRecordNotFound)) {
		if l.level >= logger.Error {
			args = append(args, "err", err)
			l.base.ErrorContext(ctx, "gorm", args...)
		}
		return
	}
	if l.slowThreshold > 0 && elapsed > l.slowThreshold {
		if l.level >= logger.Warn {
			l.base.WarnContext(ctx, "gorm slow query", args...)
		}
		return
	}
	if l.level >= logger.Info {
		l.base.InfoContext(ctx, "gorm", args...)
	}
}
