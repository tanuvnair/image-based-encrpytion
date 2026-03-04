# Architecture Overview

## Background

[Cloudflare's LavaRand](https://blog.cloudflare.com/lavarand-in-production-the-nitty-gritty-technical-details/) uses a wall of lava lamps as a source of physical entropy. A camera continuously photographs the lamps; the pixel data is unpredictable because lava lamp fluid dynamics are chaotic and impossible to reproduce exactly. The image data is hashed and fed into a CSPRNG that seeds all cryptographic operations across Cloudflare's infrastructure.

This project replicates the concept in C++:

1. **Input image** – any photograph (lava lamps, thermal noise, etc.)
2. **Entropy extraction** – pixel data is hashed to produce unbiased entropy bytes
3. **Key derivation** – a KDF (e.g. HKDF) stretches the entropy to the required key length
4. **Encryption** – a symmetric cipher (e.g. AES-256-GCM or ChaCha20-Poly1305) encrypts the payload

---

## Module Breakdown

```
include/lavaencrypt/
├── image_processor.h     # Load & pre-process the input image
├── entropy_extractor.h   # Hash pixel data → entropy byte sequence
├── encryptor.h           # Derive key from entropy; encrypt / decrypt
└── utils.h               # File I/O and hex helpers
```

### ImageProcessor

Responsible solely for I/O and pixel-level transformations.  Does **not** perform any cryptography.

### EntropyExtractor

Takes an `Image` and returns a cryptographically useful byte sequence.  The planned approach is:

1. Concatenate all pixel bytes.
2. Apply SHA-256 (or SHAKE-256 for variable output) to produce a uniform hash.
3. Optionally apply HKDF-Extract with a random salt for additional whitening.

### Encryptor

Uses the entropy bytes to:

1. Derive a symmetric key via HKDF-Expand.
2. Generate a random nonce.
3. Encrypt the payload with the chosen AEAD cipher.

The nonce and ciphertext are prepended together in the output so that the receiver can decrypt without additional metadata.

### Utils

Thin helpers: binary file I/O and hex string conversion used throughout the pipeline and in tests.

---

## Directory Layout

```
.
├── CMakeLists.txt          # Root CMake configuration
├── include/
│   └── lavaencrypt/        # Public headers (install target)
├── src/                    # Implementation + CLI entry point
├── tests/                  # Google Test unit tests
├── docs/                   # Design documentation (this file)
├── samples/                # Example input images
└── third_party/            # Vendored or notes on dependencies
```

---

## Build Instructions

```bash
cmake -B build -DLAVAENCRYPT_BUILD_TESTS=ON
cmake --build build
ctest --test-dir build --output-on-failure
```

---

## Planned Third-Party Dependencies

| Library | Purpose |
|---------|---------|
| [stb_image](https://github.com/nothings/stb) | Lightweight single-header image loader |
| [OpenSSL](https://www.openssl.org/) or [libsodium](https://libsodium.org/) | SHA-256, HKDF, AES-GCM / ChaCha20-Poly1305 |
