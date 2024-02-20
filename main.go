package main

import (
	"flag"
	"fmt"
	"github.com/stuwilli/go-versionizer/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-version"
)

type command struct {
	dir       string
	bumpLevel internal.VersionPart
}

func commandParse() *command {
	bumpLevel := flag.String("bump", "patch", "The level to bump the version (major, minor, patch)")
	flag.Parse()

	dir := flag.Arg(0) // Get the directory path
	if dir == "" {
		dir = "."
	}

	cmd, err := internal.ParseVersionPart(*bumpLevel)
	if err != nil {
		fmt.Println("Error parsing version part:", err)
		return nil
	}

	return &command{
		dir:       dir,
		bumpLevel: cmd,
	}
}

func main() {
	params := commandParse()
	// Check if the directory is a git repository
	repoPath := filepath.Join(params.dir, ".git")
	_, err := os.Stat(repoPath)
	if os.IsNotExist(err) {
		fmt.Println("No git repository found at", params.dir)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error checking for git repository:", err)
		os.Exit(1)
	}

	// Check if git is installed
	_, err = exec.LookPath("git")
	if err != nil {
		fmt.Println("git executable not found", err)
		os.Exit(1)
	}

	// Get the current git tag
	tag := readGitTag(params.dir)

	// Parse the tag as a version
	v, err := version.NewVersion(tag)
	if err != nil {
		fmt.Println("Error parsing version:", err)
		os.Exit(1)
	}

	currentVersion := internal.NewCurrentVersion(v)
	currentVersion.Bump(params.bumpLevel)
	fmt.Println(currentVersion.Version().String())
	os.Exit(0)
}

// readGitTag fetches the current git tag from the repository.
func readGitTag(dir string) string {

	cmd := exec.Command("git", "describe", "--tags", "--always", "--long")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return "0.0.0"
	}
	return strings.TrimSpace(string(out))
}
