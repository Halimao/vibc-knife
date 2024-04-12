package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

const (
	IBC_APP_SOLIDITY_TEMPLATE_REPO = "https://github.com/open-ibc/ibc-app-solidity-template"
)

func newCmd(c *cli.Context) error {
	projectName := c.Args().First()
	projectRoot := c.String("path")
	dstPath := filepath.Join(projectRoot, projectName)
	// given project root already exists
	if _, err := os.Stat(dstPath); err == nil {
		if c.Bool("force") {
			err = os.RemoveAll(dstPath)
			if err != nil {
				return err
			}
		} else {
			dstAbsPath, _ := filepath.Abs(dstPath)
			return fmt.Errorf("provided project root[%s] already exists, please check it", dstAbsPath)
		}
	}
	cloneOpts := &git.CloneOptions{
		URL:      IBC_APP_SOLIDITY_TEMPLATE_REPO,
		Progress: os.Stdout,
	}
	if c.Bool("recurse") {
		cloneOpts.RecurseSubmodules = git.DefaultSubmoduleRecursionDepth
	}

	fmt.Println("Cloning template.....")
	_, err := git.PlainClone(dstPath, false, cloneOpts)
	if err != nil {
		return err
	}
	fmt.Println("Cloned successfully")

	if err = os.RemoveAll(filepath.Join(dstPath, ".git")); err != nil {
		return err
	}
	if err = os.Rename(filepath.Join(dstPath, ".env.example"), filepath.Join(dstPath, ".env")); err != nil {
		return err
	}

	projectRootAbs, _ := filepath.Abs(dstPath)
	fmt.Printf("Create new project success!!!\nProject root: %s\n", projectRootAbs)
	return nil
}
