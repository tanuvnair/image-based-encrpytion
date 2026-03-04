#include "lavaencrypt/entropy_extractor.h"

#include <stdexcept>

namespace lavaencrypt {

std::vector<uint8_t> EntropyExtractor::extract(
    const Image& /*image*/, std::size_t /*outputBytes*/) const {
    // TODO: hash pixel data (e.g. SHA-256/SHAKE-256) and expand to outputBytes.
    throw std::runtime_error("EntropyExtractor::extract not yet implemented");
}

std::vector<uint8_t> EntropyExtractor::extractHash(const Image& /*image*/) const {
    // TODO: return SHA-256 of pixel buffer.
    throw std::runtime_error("EntropyExtractor::extractHash not yet implemented");
}

} // namespace lavaencrypt
