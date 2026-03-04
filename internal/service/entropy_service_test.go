package service_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/tanuvnair/image-based-encryption/internal/service"
)

func writeFakeImage(t *testing.T, dir, name string) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, bytes.Repeat([]byte{0xAB}, 256), 0644); err != nil {
		t.Fatalf("writeFakeImage: %v", err)
	}
}

func TestNewEntropyService_NonExistentDir(t *testing.T) {
	_, err := service.NewEntropyService("/tmp/definitely_does_not_exist_xyz")
	if err == nil {
		t.Fatal("expected error for non-existent image directory, got nil")
	}
}

func TestNewEntropyService_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	_, err := service.NewEntropyService(dir)
	if err == nil {
		t.Fatal("expected error for directory with no images, got nil")
	}
}

func TestNewEntropyService_DirWithNonImageFiles(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("hello"), 0644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	_, err := service.NewEntropyService(dir)
	if err == nil {
		t.Fatal("expected error when directory has no image files, got nil")
	}
}

func TestNewEntropyService_ValidDir(t *testing.T) {
	dir := t.TempDir()
	writeFakeImage(t, dir, "entropy01.png")
	svc, err := service.NewEntropyService(dir)
	if err != nil {
		t.Fatalf("NewEntropyService: %v", err)
	}
	if svc == nil {
		t.Fatal("NewEntropyService returned nil service without error")
	}
}

func TestEntropyService_RandomBytes_CorrectLength(t *testing.T) {
	dir := t.TempDir()
	writeFakeImage(t, dir, "entropy01.png")
	svc, err := service.NewEntropyService(dir)
	if err != nil {
		t.Fatalf("NewEntropyService: %v", err)
	}
	for _, n := range []int{1, 16, 32, 64, 256} {
		got, err := svc.RandomBytes(n)
		if err != nil {
			t.Fatalf("RandomBytes(%d): %v", n, err)
		}
		if len(got) != n {
			t.Fatalf("RandomBytes(%d) returned %d bytes", n, len(got))
		}
	}
}

func TestEntropyService_RandomBytes_NotAllZeros(t *testing.T) {
	dir := t.TempDir()
	writeFakeImage(t, dir, "entropy01.png")
	svc, _ := service.NewEntropyService(dir)
	buf, err := svc.RandomBytes(64)
	if err != nil {
		t.Fatalf("RandomBytes: %v", err)
	}
	allZero := true
	for _, b := range buf {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Fatal("RandomBytes returned all-zero output")
	}
}

func TestEntropyService_RandomBytes_TwoCallsProbablyDiffer(t *testing.T) {
	dir := t.TempDir()
	writeFakeImage(t, dir, "entropy01.png")
	svc, _ := service.NewEntropyService(dir)
	b1, err := svc.RandomBytes(64)
	if err != nil {
		t.Fatalf("first RandomBytes: %v", err)
	}
	b2, err := svc.RandomBytes(64)
	if err != nil {
		t.Fatalf("second RandomBytes: %v", err)
	}
	if bytes.Equal(b1, b2) {
		t.Fatal("two RandomBytes calls produced identical output; expected different due to fresh CSPRNG per call")
	}
}

func TestEntropyService_RandomBytes_ChiSquared(t *testing.T) {
	dir := t.TempDir()
	writeFakeImage(t, dir, "entropy01.png")
	svc, _ := service.NewEntropyService(dir)

	// Collect 100 KB of output
	var freq [256]int
	total := 100_000
	buf, _ := svc.RandomBytes(total)
	for _, b := range buf {
		freq[b]++
	}

	// Chi-squared test: expected count per bucket
	expected := float64(total) / 256.0
	var chi2 float64
	for _, count := range freq {
		diff := float64(count) - expected
		chi2 += (diff * diff) / expected
	}

	// For 255 degrees of freedom, chi2 > 310 is suspicious at p < 0.001
	if chi2 > 310 {
		t.Fatalf("chi-squared statistic %.2f suggests non-uniform output (threshold 310)", chi2)
	}
}
