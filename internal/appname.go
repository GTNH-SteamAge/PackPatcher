package internal

import (
	"fmt"
	"runtime"
	"strings"
)

const AppName = "packpatcher"

// Set via LDFLAGS -X
var (
	Version = "Unknown"
	Branch  = "Unknown"
	Commit  = "Unknown"
)

func AppVersion() string {
	return fmt.Sprintf(
		"%s version %s (git: %s@%s) (go: %s) (os: %s/%s)",
		AppName,
		Version,
		Branch,
		Commit,
		strings.TrimLeft(runtime.Version(), "go"),
		runtime.GOOS,
		runtime.GOARCH,
	)
}
