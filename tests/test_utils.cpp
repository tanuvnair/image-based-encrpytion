#include "lavaencrypt/utils.h"

#include <gtest/gtest.h>

namespace lavaencrypt::utils {

TEST(UtilsToHexString, EmptyInput) {
    EXPECT_EQ(toHexString({}), "");
}

TEST(UtilsToHexString, KnownBytes) {
    const std::vector<uint8_t> bytes = {0x00, 0xff, 0x1a, 0xb2};
    EXPECT_EQ(toHexString(bytes), "00ff1ab2");
}

TEST(UtilsFromHexString, EmptyInput) {
    EXPECT_TRUE(fromHexString("").empty());
}

TEST(UtilsFromHexString, KnownHex) {
    const auto result = fromHexString("00ff1ab2");
    const std::vector<uint8_t> expected = {0x00, 0xff, 0x1a, 0xb2};
    EXPECT_EQ(result, expected);
}

TEST(UtilsFromHexString, OddLengthThrows) {
    EXPECT_THROW(fromHexString("abc"), std::invalid_argument);
}

TEST(UtilsFromHexString, InvalidCharThrows) {
    EXPECT_THROW(fromHexString("zz"), std::invalid_argument);
}

TEST(UtilsRoundTrip, HexRoundTrip) {
    const std::vector<uint8_t> original = {0xde, 0xad, 0xbe, 0xef};
    EXPECT_EQ(fromHexString(toHexString(original)), original);
}

} // namespace lavaencrypt::utils
