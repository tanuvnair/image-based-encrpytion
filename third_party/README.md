# Third-Party Dependencies

This directory is reserved for vendored third-party code or notes on external dependencies.

## Planned Dependencies

### stb_image (image loading)
- **Source**: https://github.com/nothings/stb
- **License**: MIT / Public Domain
- **Usage**: Single-header library for loading PNG/JPEG/BMP/etc. images
- **Integration**: Copy `stb_image.h` into `third_party/stb/` and add a CMake interface target

### OpenSSL or libsodium (cryptography)
- **OpenSSL source**: https://www.openssl.org/
- **libsodium source**: https://libsodium.org/
- **Usage**: SHA-256, HKDF, AES-256-GCM, ChaCha20-Poly1305
- **Integration**: Locate via `find_package(OpenSSL)` or `find_package(sodium)` in CMake

## Adding a Vendored Library

1. Place the source under `third_party/<library-name>/`.
2. Add a `CMakeLists.txt` (or use `add_subdirectory`) to expose it as a CMake target.
3. Link it from `src/CMakeLists.txt` via `target_link_libraries`.
