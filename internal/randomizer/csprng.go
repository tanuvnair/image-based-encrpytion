package randomizer

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// CSPRNG is a cryptographically secure pseudo-random number generator
// seeded by the Mixer output. It uses AES-CTR as the underlying stream
// cipher — available in Go's stdlib with no external dependencies.
type CSPRNG struct {
	stream cipher.Stream
}

// NewCSPRNG creates a new CSPRNG seeded by the provided seed bytes.
// The seed must be at least 48 bytes (32 for AES-256 key + 16 for IV).
// The Mixer's SHA-512 output (64 bytes) satisfies this requirement.
func NewCSPRNG(seed []byte) (*CSPRNG, error) {
	if len(seed) < 48 {
		return nil, fmt.Errorf("csprng: seed must be at least 48 bytes, got %d", len(seed))
	}

	// First 32 bytes → AES-256 key
	// Next  16 bytes → AES block size IV (counter start)
	key := seed[:32]
	iv := seed[32:48]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("csprng: creating AES cipher: %w", err)
	}

	stream := cipher.NewCTR(block, iv)
	return &CSPRNG{stream: stream}, nil
}

// Read fills p with pseudo-random bytes, satisfying io.Reader.
// XORs a zero-value buffer with the AES-CTR keystream to extract
// raw random bytes.
func (c *CSPRNG) Read(p []byte) (int, error) {
	zeros := make([]byte, len(p))
	c.stream.XORKeyStream(p, zeros)
	return len(p), nil
}
