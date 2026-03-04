#pragma once

#include "image_processor.h"

#include <cstdint>
#include <vector>

namespace lavaencrypt {

/// Extracts high-quality entropy from the chaotic visual state captured in an
/// image (e.g. a lava-lamp photo as used by Cloudflare's LavaRand system).
///
/// The extractor hashes the pixel data to produce an unpredictable byte
/// sequence that is suitable for seeding a cryptographically secure random
/// number generator (CSPRNG).
class EntropyExtractor {
public:
    EntropyExtractor() = default;
    ~EntropyExtractor() = default;

    /// Extract raw entropy bytes from @p image.
    /// @return A vector of entropy bytes whose length equals @p outputBytes.
    std::vector<uint8_t> extract(const Image& image, std::size_t outputBytes) const;

    /// Convenience overload: extract the full SHA-256 hash (32 bytes) of the
    /// pixel buffer as a quick entropy source.
    std::vector<uint8_t> extractHash(const Image& image) const;
};

} // namespace lavaencrypt
