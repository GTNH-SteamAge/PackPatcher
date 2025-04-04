package git

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v70/github"
)

var (
	ctx = context.Background()
	org = "GTNH-SteamAge"
)

func (c *Client) GetLatestRelease(repoName string) (*github.RepositoryRelease, error) {
	repo, _, err := c.C().Repositories.GetLatestRelease(ctx, org, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest release for repo %s: %w", repoName, err)
	}
	return repo, nil
}

func (c *Client) GetObfJarAsset(release *github.RepositoryRelease) (*github.ReleaseAsset, error) {
	for _, asset := range release.Assets {
		assetName := asset.GetName()
		if !strings.HasSuffix(assetName, ".jar") {
			continue
		}
		if strings.HasSuffix(assetName, "-dev.jar") {
			continue
		}
		if strings.HasSuffix(assetName, "-sources.jar") {
			continue
		}
		if strings.HasSuffix(assetName, "-dev-preshadow.jar") {
			continue
		}
		return asset, nil
	}
	return nil, fmt.Errorf("failed to find obf jar asset for release %s", release.GetName())
}
