package tools

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Goimports runs goimports on the provided directory files.
func Goimports(dir string) error {
	if _, err := exec.LookPath("goimports"); err != nil {
		return fmt.Errorf("goimports not found: %w", err)
	}
	args := []string{"-w"}
	err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(d.Name(), ".go") {
			args = append(args, filepath.Join(dir, d.Name()))
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walkdir: %w", err)
	}
	cmd := exec.Command("goimports", args...)
	cmd.Stderr = cmd.Stdout
	if stdout, err := cmd.Output(); err != nil {
		return fmt.Errorf("goimports: %w: %s", err, string(stdout))
	}
	return nil
}

// GitStatus runs git status on the provided directory.
func GitStatus(dir string) ([]string, error) {
	if _, err := exec.LookPath("git"); err != nil {
		return nil, fmt.Errorf("git not found: %w", err)
	}
	cmd := exec.Command("git", "status", "--porcelain", dir)
	cmd.Stderr = cmd.Stdout
	stdout, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git status: %w: %s", err, string(stdout))
	}
	var diff []string
	for _, line := range strings.Split(string(stdout), "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			diff = append(diff, line)
		}
	}
	return diff, nil
}
