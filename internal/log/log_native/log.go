package lognative

import (
	"context"
	"log/slog"
	"os"
)

type LogNative struct {
	logger *slog.Logger
}

func NewLogNative() *LogNative {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &LogNative{logger: logger}
}

func (ref *LogNative) Info(ctx context.Context, msg, keyArgs string, args ...any) {
	ref.logger.InfoContext(ctx, msg, keyArgs, args)
}

func (ref *LogNative) Warn(ctx context.Context, msg, keyArgs string, args ...any) {
	ref.logger.WarnContext(ctx, msg, keyArgs, args)
}

func (ref *LogNative) Debug(ctx context.Context, msg, keyArgs string, args ...any) {
	ref.logger.DebugContext(ctx, msg, keyArgs, args)
}

func (ref *LogNative) Error(ctx context.Context, msg, keyArgs string, err error) error {
	ref.logger.ErrorContext(ctx, msg, keyArgs, err.Error())
	return err
}
