package log

import "context"

type Logger interface {
	Info(ctx context.Context, msg, keyArgs string, args ...any)
	Warn(ctx context.Context, msg, keyArgs string, args ...any)
	Debug(ctx context.Context, msg, keyArgs string, args ...any)
	Error(ctx context.Context, msg, keyArgs string, err error) error
}
