package utils_test

import (
"os"
"path/filepath"
"testing"

"github.com/tanuvnair/image-based-encryption/internal/utils"
)

func TestValidateImageDir_ValidDirectory(t *testing.T) {
	dir := t.TempDir()
	if err := utils.ValidateImageDir(dir); err != nil {
		t.Fatalf("ValidateImageDir(%q) returned unexpected error: %v", dir, err)
	}
}

func TestValidateImageDir_NonExistentPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does_not_exist")
	if err := utils.ValidateImageDir(path); err == nil {
		t.Fatalf("expected error for non-existent path %q, got nil", path)
	}
}

func TestValidateImageDir_FileNotDirectory(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(file, []byte("data"), 0644); err != nil {
		t.Fatalf("creating test file: %v", err)
	}
	if err := utils.ValidateImageDir(file); err == nil {
		t.Fatal("expected error when path is a file, got nil")
	}
}
