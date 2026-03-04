# image-based-encryption

A C++ replication of [Cloudflare's LavaRand](https://blog.cloudflare.com/lavarand-in-production-the-nitty-gritty-technical-details/) — image-based entropy extraction and encryption.

## Concept

Cloudflare photographs a wall of lava lamps whose chaotic fluid dynamics produce unpredictable pixel data. That data is hashed to seed a cryptographically secure random number generator. This project implements the same pipeline: load an image → extract entropy → derive a symmetric key → encrypt a payload.

## Project Structure

```
.
├── CMakeLists.txt          # Root CMake build file
├── include/
│   └── lavaencrypt/        # Public C++ headers
│       ├── image_processor.h
│       ├── entropy_extractor.h
│       ├── encryptor.h
│       └── utils.h
├── src/                    # Implementation files + CLI entry point
├── tests/                  # Google Test unit tests
├── docs/
│   └── architecture.md     # Design and architecture notes
├── samples/                # Example input images
└── third_party/            # Vendored or noted external dependencies
```

## Build

```bash
# Configure (tests enabled by default)
cmake -B build -DCMAKE_BUILD_TYPE=Release

# Build
cmake --build build

# Run tests
ctest --test-dir build --output-on-failure
```

Requires CMake ≥ 3.20 and a C++17-capable compiler.  
Google Test is fetched automatically by CMake's `FetchContent`.

## Usage

```
lavaencrypt <input_image> <plaintext_file> <output_file>
```

> **Note**: The encryption pipeline is not yet implemented. See `docs/architecture.md` for the planned design.

## License

MIT
