#include "lavaencrypt/entropy_extractor.h"

#include <gtest/gtest.h>

namespace lavaencrypt {

TEST(EntropyExtractorExtract, NotYetImplementedThrows) {
    const EntropyExtractor extractor;
    const Image dummy;
    EXPECT_THROW(extractor.extract(dummy, 32), std::exception);
}

TEST(EntropyExtractorExtractHash, NotYetImplementedThrows) {
    const EntropyExtractor extractor;
    const Image dummy;
    EXPECT_THROW(extractor.extractHash(dummy), std::exception);
}

} // namespace lavaencrypt
