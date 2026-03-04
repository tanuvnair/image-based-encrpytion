package service

import (
	"fmt"

	"github.com/tanuvnair/image-based-encryption/internal/entropy"
	"github.com/tanuvnair/image-based-encryption/internal/randomizer"
)

// EntropyService wires together the ImageSource, Mixer and CSPRNG.
// Each call to RandomBytes re-seeds the CSPRNG from a fresh image pick,
// so every request gets a unique entropy source.
type EntropyService struct {
	mixer *entropy.Mixer
}

// NewEntropyService creates an EntropyService using images from imageDir.
func NewEntropyService(imageDir string) (*EntropyService, error) {
	imgSrc, err := entropy.NewImageSource(imageDir)
	if err != nil {
		return nil, fmt.Errorf("entropy service: creating image source: %w", err)
	}

	mixer := entropy.NewMixer(imgSrc)
	return &EntropyService{mixer: mixer}, nil
}

// RandomBytes generates n cryptographically random bytes by:
//  1. Picking a fresh random image and mixing it with crypto/rand → seed
//  2. Seeding a new AES-CTR CSPRNG with that seed
//  3. Reading n bytes from the CSPRNG
func (s *EntropyService) RandomBytes(n int) ([]byte, error) {
	seed, err := s.mixer.Seed()
	if err != nil {
		return nil, fmt.Errorf("entropy service: generating seed: %w", err)
	}

	csprng, err := randomizer.NewCSPRNG(seed)
	if err != nil {
		return nil, fmt.Errorf("entropy service: creating CSPRNG: %w", err)
	}

	buf := make([]byte, n)
	if _, err := csprng.Read(buf); err != nil {
		return nil, fmt.Errorf("entropy service: reading random bytes: %w", err)
	}

	return buf, nil
}
