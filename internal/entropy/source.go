package entropy

// Source provides raw unpredictable bytes from a physical or digital source.
// Implementations might capture camera frames, read sensor data, or load
// image files from disk.
type Source interface {
	Entropy() ([]byte, error)
}
