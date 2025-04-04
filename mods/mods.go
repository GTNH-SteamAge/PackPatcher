package mods

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/v70/github"

	"github.com/GTNH-SteamAge/PackPatcher/git"
)

var (
	repositoryNames = []string{
		"GT5-Unofficial",
	}
)

func GatherMods(archiveDir string) error {
	client, err := git.GetClient("")
	if err != nil {
		return fmt.Errorf("failed to get github client: %w", err)
	}

	releases := make([]*github.RepositoryRelease, 0)
	for _, repo := range repositoryNames {
		release, err := client.GetLatestRelease(repo)
		if err != nil {
			return err
		}
		releases = append(releases, release)
	}

	modsPath := filepath.Join(archiveDir, "mods")
	if err = os.MkdirAll(modsPath, 0750); err != nil {
		return fmt.Errorf("failed to create mods tmp dir: %w", err)
	}

	for _, release := range releases {
		asset, err := client.GetObfJarAsset(release)
		if err != nil {
			return err
		}

		url := asset.GetBrowserDownloadURL()
		if err := client.DownloadFile(url, filepath.Join(modsPath, release.GetName())); err != nil {
			return err
		}
	}

	return nil
}
