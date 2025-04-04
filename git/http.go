package git

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func (c *Client) DownloadFile(url string, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	zap.S().Infof("downloaded file %s to %s", url, dest)
	return nil
}
