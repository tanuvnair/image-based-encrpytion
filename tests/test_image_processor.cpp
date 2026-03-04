#include "lavaencrypt/image_processor.h"

#include <gtest/gtest.h>

namespace lavaencrypt {

TEST(ImageProcessorLoad, NonExistentFileThrows) {
    const ImageProcessor processor;
    EXPECT_THROW(processor.load("/nonexistent/path/image.png"), std::exception);
}

TEST(ImageProcessorResize, NotYetImplementedThrows) {
    const ImageProcessor processor;
    const Image dummy;
    EXPECT_THROW(processor.resize(dummy, 64, 64), std::exception);
}

TEST(ImageProcessorGreyscale, NotYetImplementedThrows) {
    const ImageProcessor processor;
    const Image dummy;
    EXPECT_THROW(processor.toGreyscale(dummy), std::exception);
}

} // namespace lavaencrypt
