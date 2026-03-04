package entropy

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

// supportedExtensions lists file extensions recognized as images.
var supportedExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".bmp":  true,
	".webp": true,
	".tiff": true,
}

// ImageSource reads a random image from a directory and returns its raw bytes.
type ImageSource struct {
	paths []string
}

// NewImageSource scans the directory for image files and returns a ready-to-use
// ImageSource. It returns an error if no supported image files are found.
func NewImageSource(dir string) (*ImageSource, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %q: %w", dir, err)
	}

	var paths []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if supportedExtensions[ext] {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("no supported image files found in %q", dir)
	}

	return &ImageSource{paths: paths}, nil
}

// Entropy picks a random image from the directory and returns its raw bytes.
// The random selection uses crypto/rand to avoid predictability in which image
// is chosen.
func (s *ImageSource) Entropy() ([]byte, error) {
	idx, err := cryptoRandIndex(len(s.paths))
	if err != nil {
		return nil, fmt.Errorf("selecting random image: %w", err)
	}

	data, err := os.ReadFile(s.paths[idx])
	if err != nil {
		return nil, fmt.Errorf("reading image %q: %w", s.paths[idx], err)
	}

	return data, nil
}

// cryptoRandIndex returns a cryptographically random integer in [0, n).
func cryptoRandIndex(n int) (int, error) {
	max := big.NewInt(int64(n))
	idx, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return int(idx.Int64()), nil
}
