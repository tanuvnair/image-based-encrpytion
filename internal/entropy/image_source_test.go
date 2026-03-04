package entropy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tanuvnair/image-based-encryption/internal/entropy"
)

// createTestImage writes a small valid PNG-like file to the given directory.
// We don't need a real PNG -- we just need bytes on disk.
func createTestImage(t *testing.T, dir, name string, content []byte) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatalf("failed to create test image %s: %v", path, err)
	}
}

func TestSource_Entropy_ReturnsBytesFromAnImage(t *testing.T) {
	dir := t.TempDir()
	content := []byte("fake image data with some bytes")
	createTestImage(t, dir, "lamp01.png", content)

	src, err := entropy.NewImageSource(dir)
	if err != nil {
		t.Fatalf("NewImageSource(%q) returned error: %v", dir, err)
	}

	got, err := src.Entropy()
	if err != nil {
		t.Fatalf("Entropy() returned error: %v", err)
	}
	if len(got) == 0 {
		t.Fatal("Entropy() returned zero bytes")
	}
}

func TestSource_NewImageSource_FailsOnEmptyDir(t *testing.T) {
	dir := t.TempDir()
	// No images in the directory.
	_, err := entropy.NewImageSource(dir)
	if err == nil {
		t.Fatal("expected error for empty directory, got nil")
	}
}

func TestSource_NewImageSource_IgnoresNonImageFiles(t *testing.T) {
	dir := t.TempDir()
	createTestImage(t, dir, "notes.txt", []byte("not an image"))
	_, err := entropy.NewImageSource(dir)
	if err == nil {
		t.Fatal("expected error when directory has no image files, got nil")
	}
}
