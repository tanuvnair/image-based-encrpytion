package entropy

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
)

// Mixer combines entropy from a Source (image bytes) with entropy from the
// OS via crypto/rand. The two streams are XOR-ed together and then hashed
// with SHA-512 to produce a fixed-size 64-byte seed.
//
// This mirrors Cloudflare's LavaRand approach: even if the image entropy is
// somehow compromised, the crypto/rand layer keeps the output unpredictable,
// and vice versa.
type Mixer struct {
	source Source
}

// NewMixer returns a Mixer that draws physical entropy from src.
func NewMixer(src Source) *Mixer {
	return &Mixer{source: src}
}

// Seed produces a 64-byte cryptographic seed by:
//  1. Reading raw bytes from the entropy Source (image binary).
//  2. Reading an equal number of bytes from crypto/rand.
//  3. XOR-ing the two streams together.
//  4. Hashing the result with SHA-512.
func (m *Mixer) Seed() ([]byte, error) {
	imgBytes, err := m.source.Entropy()
	if err != nil {
		return nil, fmt.Errorf("mixer: reading image entropy: %w", err)
	}

	if len(imgBytes) == 0 {
		return nil, fmt.Errorf("mixer: image source returned empty bytes")
	}

	randBytes := make([]byte, len(imgBytes))
	if _, err := rand.Read(randBytes); err != nil {
		return nil, fmt.Errorf("mixer: reading crypto/rand bytes: %w", err)
	}

	// XOR the image bytes with the crypto/rand bytes.
	mixed := make([]byte, len(imgBytes))
	for i := range imgBytes {
		mixed[i] = imgBytes[i] ^ randBytes[i]
	}

	// Collapse the mixed stream into a fixed-size seed via SHA-512.
	hash := sha512.Sum512(mixed)
	return hash[:], nil
}
