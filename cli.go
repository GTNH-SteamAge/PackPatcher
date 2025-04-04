package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/GTNH-SteamAge/PackPatcher/internal"
)

const cliDescription = `` // todo

var debugLogs bool

func main() {
	app := &cli.App{
		Name:        internal.AppName,
		Usage:       "todo",
		Description: cliDescription,
		Suggest:     true,
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func init() {
	zap.ReplaceGlobals(zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		}),
		zapcore.Lock(os.Stderr),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return debugLogs || lvl > zapcore.DebugLevel
		}))))
}
