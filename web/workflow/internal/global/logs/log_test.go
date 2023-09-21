package logs

import (
	"context"
	"envoy-go-fliter-hub/config"
	"go.uber.org/zap"

	"os"
	"path/filepath"
	"testing"
)

func TestLogFile(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	dir = filepath.Clean(filepath.Join(dir, "logs"))
	config.Config.Log.FilePath = dir
	config.Config.Mode = config.ReleaseMode

	_ = os.RemoveAll(dir)
	Init()

	Debug("debug")
	Info("info", zap.Any("key", "any"))
	Warn("warn")
	Error("error")

	Debugf("debugf")
	Infof("infof, val=%v", "val")
	Warnf("warnf")
	Errorf("errorf")

	ctx := CtxAddKVs(context.TODO(), zap.Int("log_id", 123))
	CtxDebug(ctx, "ctx debug")
	CtxInfo(ctx, "ctx info", zap.Any("key", "any"))
	CtxWarn(ctx, "ctx warn")
	CtxError(ctx, "ctx error")
}

func TestConsole(t *testing.T) {
	config.Config.Mode = config.DebugMode
	Init()

	Debug("debug")
	Info("info", zap.Any("key", "any"))
	Warn("warn")
	Error("error")

	Debugf("debugf")
	Infof("infof, val=%v", "val")
	Warnf("warnf")
	Errorf("errorf")

	ctx := CtxAddKVs(context.TODO(), zap.Int("log_id", 123))
	CtxDebug(ctx, "ctx debug")
	CtxInfo(ctx, "ctx info", zap.Any("key", "any"))
	CtxWarn(ctx, "ctx warn")
	CtxError(ctx, "ctx error")
}
