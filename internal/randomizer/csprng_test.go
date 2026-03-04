package randomizer_test

import (
	"bytes"
	"testing"

	"github.com/tanuvnair/image-based-encryption/internal/randomizer"
)

// validSeed returns a 64-byte seed suitable for NewCSPRNG.
func validSeed() []byte {
	seed := make([]byte, 64)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	return seed
}

func TestNewCSPRNG_ValidSeed(t *testing.T) {
	_, err := randomizer.NewCSPRNG(validSeed())
	if err != nil {
		t.Fatalf("NewCSPRNG with valid seed returned error: %v", err)
	}
}

func TestNewCSPRNG_ShortSeed(t *testing.T) {
	shortSeed := make([]byte, 47)
	_, err := randomizer.NewCSPRNG(shortSeed)
	if err == nil {
		t.Fatal("expected error for seed shorter than 48 bytes, got nil")
	}
}

func TestNewCSPRNG_ExactMinimumSeed(t *testing.T) {
	seed := make([]byte, 48)
	for i := range seed {
		seed[i] = byte(i + 1)
	}
	_, err := randomizer.NewCSPRNG(seed)
	if err != nil {
		t.Fatalf("NewCSPRNG with 48-byte seed returned error: %v", err)
	}
}

func TestCSPRNG_Read_ProducesCorrectLength(t *testing.T) {
	csprng, err := randomizer.NewCSPRNG(validSeed())
	if err != nil {
		t.Fatalf("NewCSPRNG: %v", err)
	}

	for _, n := range []int{1, 16, 32, 64, 256} {
		buf := make([]byte, n)
		got, err := csprng.Read(buf)
		if err != nil {
			t.Fatalf("Read(%d): unexpected error: %v", n, err)
		}
		if got != n {
			t.Fatalf("Read(%d): returned %d bytes, want %d", n, got, n)
		}
	}
}

func TestCSPRNG_Read_OutputNotAllZeros(t *testing.T) {
	csprng, err := randomizer.NewCSPRNG(validSeed())
	if err != nil {
		t.Fatalf("NewCSPRNG: %v", err)
	}

	buf := make([]byte, 64)
	if _, err := csprng.Read(buf); err != nil {
		t.Fatalf("Read: %v", err)
	}

	allZero := true
	for _, b := range buf {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Fatal("Read produced all-zero output, expected pseudo-random bytes")
	}
}

func TestCSPRNG_Read_DeterministicWithSameSeed(t *testing.T) {
	seed := validSeed()

	csprng1, _ := randomizer.NewCSPRNG(seed)
	csprng2, _ := randomizer.NewCSPRNG(seed)

	buf1 := make([]byte, 64)
	buf2 := make([]byte, 64)
	csprng1.Read(buf1) //nolint:errcheck
	csprng2.Read(buf2) //nolint:errcheck

	if !bytes.Equal(buf1, buf2) {
		t.Fatal("two CSPRNGs seeded identically produced different output")
	}
}

func TestCSPRNG_Read_DifferentSeeds_DifferentOutput(t *testing.T) {
	seed1 := validSeed()
	seed2 := validSeed()
	seed2[0] ^= 0xFF // flip some bits

	csprng1, _ := randomizer.NewCSPRNG(seed1)
	csprng2, _ := randomizer.NewCSPRNG(seed2)

	buf1 := make([]byte, 64)
	buf2 := make([]byte, 64)
	csprng1.Read(buf1) //nolint:errcheck
	csprng2.Read(buf2) //nolint:errcheck

	if bytes.Equal(buf1, buf2) {
		t.Fatal("two CSPRNGs with different seeds produced identical output")
	}
}
