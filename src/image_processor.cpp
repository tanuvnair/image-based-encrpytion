#include "lavaencrypt/image_processor.h"

#include <stdexcept>

namespace lavaencrypt {

Image ImageProcessor::load(const std::string& /*filepath*/) const {
    // TODO: use a library such as stb_image or OpenCV to load the image.
    throw std::runtime_error("ImageProcessor::load not yet implemented");
}

Image ImageProcessor::resize(const Image& /*src*/, int /*width*/, int /*height*/) const {
    // TODO: implement image resizing.
    throw std::runtime_error("ImageProcessor::resize not yet implemented");
}

Image ImageProcessor::toGreyscale(const Image& /*src*/) const {
    // TODO: implement greyscale conversion.
    throw std::runtime_error("ImageProcessor::toGreyscale not yet implemented");
}

} // namespace lavaencrypt
