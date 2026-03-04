#pragma once

#include <cstdint>
#include <string>
#include <vector>

namespace lavaencrypt {

/// Represents a raw image loaded into memory as a flat RGBA pixel buffer.
struct Image {
    std::vector<uint8_t> pixels; ///< Raw pixel data (RGBA, row-major)
    int width  = 0;              ///< Image width in pixels
    int height = 0;              ///< Image height in pixels
    int channels = 0;            ///< Number of colour channels (e.g. 3=RGB, 4=RGBA)
};

/// Loads and pre-processes the input image so that it can be passed to the
/// EntropyExtractor.  No encryption logic lives here.
class ImageProcessor {
public:
    ImageProcessor() = default;
    ~ImageProcessor() = default;

    /// Load an image from @p filepath into an Image struct.
    /// @throws std::runtime_error if the file cannot be opened or decoded.
    Image load(const std::string& filepath) const;

    /// Optionally resize the image to @p width x @p height before processing.
    Image resize(const Image& src, int width, int height) const;

    /// Convert the image to greyscale in-place (reduces channel count to 1).
    Image toGreyscale(const Image& src) const;
};

} // namespace lavaencrypt
