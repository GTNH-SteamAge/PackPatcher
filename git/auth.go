package git

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/google/go-github/v70/github"
	"golang.org/x/oauth2"
)

type Client struct {
	c *github.Client
}

func (c *Client) C() *github.Client {
	return c.c
}

func GetClient(token string) (*Client, error) {
	token = getToken(token)
	if token == "" {
		return nil, errors.New("could not find github token")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Client{c: client}, nil
}

func getToken(token string) string {
	if token != "" {
		return token
	}

	token = os.Getenv("GITHUB_TOKEN")
	if token != "" {
		return token
	}

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "windows":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		fi, err := os.Open(filepath.Join(homeDir, ".github_personal_token"))
		if err != nil {
			return ""
		}
		data, err := io.ReadAll(fi)
		if err != nil {
			return ""
		}
		return string(data)
	case "linux":
		fi, err := os.Open(filepath.Join("/", ".github_personal_token"))
		if err != nil {
			return ""
		}
		data, err := io.ReadAll(fi)
		if err != nil {
			return ""
		}
		return string(data)
	default:
		return ""
	}
}
