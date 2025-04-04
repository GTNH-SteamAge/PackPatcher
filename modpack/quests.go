package modpack

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GTNH-SteamAge/PackPatcher/git"
)

func GatherQuestFiles(archiveDir string) error {
	tempDir, err := os.MkdirTemp("", "GT-New-Horizons-Modpack")
	if err != nil {
		return fmt.Errorf("failed to create tmp dir: %w", err)
	}
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	_, err = git.CloneRepo(tempDir, "GT-New-Horizons-Modpack", "steam-age")
	if err != nil {
		return fmt.Errorf("failed to clone modpack repo: %w", err)
	}

	questsPath := filepath.Join("config", "betterquesting", "DefaultQuests")
	questsPathSrc := filepath.Join(tempDir, questsPath)
	questsPathDst := filepath.Join(archiveDir, questsPath)

	if err = os.MkdirAll(questsPathDst, 0750); err != nil {
		return fmt.Errorf("failed to create quests tmp dir: %w", err)
	}

	if err = os.CopyFS(questsPathDst, os.DirFS(questsPathSrc)); err != nil {
		return fmt.Errorf("failed to copy quests to tmp dir: %w", err)
	}

	return nil
}
