package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/mod/semver"

	"github.com/GTNH-SteamAge/PackPatcher/modpack"
)

var debugLogs bool

func run() bool {
	for _, str := range os.Args {
		if strings.Contains(str, "--debug") {
			debugLogs = true
			break
		}
	}

	if len(os.Args) < 2 {
		zap.S().Error("Please provide a semantic version (e.g. v1.2.3)")
		return false
	}

	version := os.Args[1]
	if !semver.IsValid(version) {
		zap.S().Errorf("Provided release (%s) not in semantic format (e.g. v1.2.3)", version)
		return false
	}

	tempDir, err := os.MkdirTemp("", fmt.Sprintf("GTNH-SteamAgePatch-%s", version))
	if err != nil {
		zap.S().Errorf("Failed to create temporary directory: %v", err)
		return false
	}
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	if err := modpack.GatherQuestFiles(tempDir); err != nil {
		zap.S().Errorf("Failed to gather modpack quest files: %v", err)
		return false
	}

	//if err := mods.GatherMods(tempDir); err != nil {
	//	zap.S().Errorf("Failed to gather mod jars: %v", err)
	//	return false
	//}

	// testing code
	_ = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.Size())
		return nil
	})

	return true
}

func main() {
	if !run() {
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
