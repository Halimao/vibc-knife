package downloader

import (
	"context"
	"os"

	getter "github.com/hashicorp/go-getter/v2"
)

func DownloadRepo(repoUrl, dstPath string) error {
	// Get the pwd
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Build the client
	req := &getter.Request{
		Src:              repoUrl,
		Dst:              dstPath,
		Pwd:              pwd,
		ProgressListener: defaultProgressBar,
	}
	_, err = getter.DefaultClient.Get(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
