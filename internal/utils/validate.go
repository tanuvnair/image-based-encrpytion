package utils

import (
	"fmt"
	"os"
)

// ValidateImageDir checks that the given path exists and is a directory.
func ValidateImageDir(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("cannot access %q: %w", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%q is not a directory", path)
	}
	return nil
}
