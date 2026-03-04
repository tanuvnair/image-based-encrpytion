package handlers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tanuvnair/image-based-encryption/internal/handlers"
)

// fakeSource implements handlers.RandomSource for testing.
type fakeSource struct {
	data []byte
	err  error
}

func (f *fakeSource) RandomBytes(n int) ([]byte, error) {
	if f.err != nil {
		return nil, f.err
	}
	// Return a deterministic slice of n bytes.
	out := make([]byte, n)
	for i := range out {
		out[i] = byte(i % 256)
	}
	return out, nil
}

func newRequest(query string) *http.Request {
	url := "/random"
	if query != "" {
		url = fmt.Sprintf("/random?%s", query)
	}
	req := httptest.NewRequest(http.MethodGet, url, nil)
	return req
}

func TestHandleRandom_DefaultBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest(""))

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	// 32 bytes = 64 hex chars
	if !strings.Contains(body, "random bytes (32)") {
		t.Errorf("unexpected body: %q", body)
	}
}

func TestHandleRandom_CustomBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=16"))

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "random bytes (16)") {
		t.Errorf("unexpected body: %q", w.Body.String())
	}
}

func TestHandleRandom_MaxBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=1024"))

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestHandleRandom_ExceedsMaxBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=1025"))

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleRandom_InvalidBytesParam(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=abc"))

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleRandom_NegativeBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=-1"))

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestHandleRandom_ZeroBytes(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest("bytes=0"))

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for zero bytes, got %d", w.Code)
	}
}

func TestHandleRandom_SourceError(t *testing.T) {
	src := &fakeSource{err: errors.New("source failure")}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest(""))

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestHandleRandom_ContentTypeIsTextPlain(t *testing.T) {
	src := &fakeSource{}
	handler := handlers.HandleRandom(src)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, newRequest(""))

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "text/plain") {
		t.Fatalf("expected Content-Type text/plain, got %q", ct)
	}
}
