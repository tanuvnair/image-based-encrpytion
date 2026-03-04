package entropy_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/tanuvnair/image-based-encryption/internal/entropy"
)

// staticSource implements entropy.Source and always returns the same bytes.
type staticSource struct {
	data []byte
	err  error
}

func (s *staticSource) Entropy() ([]byte, error) {
	return s.data, s.err
}

func TestMixer_Seed_Returns64Bytes(t *testing.T) {
	src := &staticSource{data: bytes.Repeat([]byte{0xAB}, 128)}
	mixer := entropy.NewMixer(src)

	seed, err := mixer.Seed()
	if err != nil {
		t.Fatalf("Seed() returned error: %v", err)
	}
	if len(seed) != 64 {
		t.Fatalf("Seed() returned %d bytes, want 64 (SHA-512 output)", len(seed))
	}
}

func TestMixer_Seed_SourceError_Propagates(t *testing.T) {
	src := &staticSource{err: errors.New("disk error")}
	mixer := entropy.NewMixer(src)

	_, err := mixer.Seed()
	if err == nil {
		t.Fatal("expected error when source fails, got nil")
	}
}

func TestMixer_Seed_EmptySourceBytes_ReturnsError(t *testing.T) {
	src := &staticSource{data: []byte{}}
	mixer := entropy.NewMixer(src)

	_, err := mixer.Seed()
	if err == nil {
		t.Fatal("expected error when source returns empty bytes, got nil")
	}
}

func TestMixer_Seed_TwoCallsProbablyDiffer(t *testing.T) {
	// Because crypto/rand contributes entropy each call the seeds should differ.
	// There is an astronomically small (2^-512) chance this could collide.
	src := &staticSource{data: bytes.Repeat([]byte{0x01}, 64)}
	mixer := entropy.NewMixer(src)

	seed1, err := mixer.Seed()
	if err != nil {
		t.Fatalf("first Seed(): %v", err)
	}
	seed2, err := mixer.Seed()
	if err != nil {
		t.Fatalf("second Seed(): %v", err)
	}

	if bytes.Equal(seed1, seed2) {
		t.Fatal("two Seed() calls produced identical output; crypto/rand should make this unique")
	}
}

func TestMixer_Seed_OutputNot64ZeroBytes(t *testing.T) {
	src := &staticSource{data: bytes.Repeat([]byte{0xFF}, 64)}
	mixer := entropy.NewMixer(src)

	seed, err := mixer.Seed()
	if err != nil {
		t.Fatalf("Seed(): %v", err)
	}

	allZero := true
	for _, b := range seed {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Fatal("Seed() returned all-zero bytes")
	}
}
