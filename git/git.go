package git

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"go.uber.org/zap"
)

func CloneRepo(clonePath string, repoName string, branchName string) (*git.Repository, error) {
	zap.S().Infof("Cloning https://github.com/GTNH-SteamAge/%s...", repoName)
	repo, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		Tags:         git.NoTags,
		URL:          fmt.Sprintf("git@github.com:GTNH-SteamAge/%s.git", repoName),
		SingleBranch: false,
		Depth:        1,
	})
	if err == nil {
		return repo, checkoutBranch(repo, branchName)
	}

	// If git exited with an error, try to use HTTP instead
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		repo, err = git.PlainClone(clonePath, false, &git.CloneOptions{
			Tags:         git.NoTags,
			URL:          fmt.Sprintf("https://github.com/GTNH-SteamAge/%s.git", repoName),
			SingleBranch: false,
			Depth:        1,
		})

		if err != nil {
			return nil, err
		}
	}
	return repo, nil
}

func checkoutBranch(repo *git.Repository, branchName string) error {
	tree, err := repo.Worktree()
	if err != nil {
		return err
	}

	branchRefName := plumbing.NewBranchReferenceName(branchName)
	branchOpts := git.CheckoutOptions{
		Branch: branchRefName,
		Force:  true,
	}
	if err := tree.Checkout(&branchOpts); err != nil {
		mirrorRemoteBranchRefSpec := fmt.Sprintf("refs/heads/%s:refs/heads/%s", branchName, branchName)
		if err := fetchOrigin(repo, mirrorRemoteBranchRefSpec); err != nil {
			return err
		}

		if err := tree.Checkout(&branchOpts); err != nil {
			return err
		}
	}

	zap.S().Info("git show-ref --head HEAD")
	ref, err := repo.Head()
	if err != nil {
		return err
	}
	zap.S().Info(ref.Hash())
	return nil
}

func fetchOrigin(repo *git.Repository, refSpecStr string) error {
	remote, err := repo.Remote("origin")
	if err != nil {
		return err
	}

	var refSpecs []config.RefSpec
	if refSpecStr != "" {
		refSpecs = []config.RefSpec{config.RefSpec(refSpecStr)}
	}

	if err := remote.Fetch(&git.FetchOptions{
		RefSpecs: refSpecs,
	}); err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			zap.S().Info("refs already up to date")
		} else {
			return fmt.Errorf("fetch origin failed: %w", err)
		}
	}

	return nil
}
