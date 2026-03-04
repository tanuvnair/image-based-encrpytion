#include "lavaencrypt/encryptor.h"

#include <gtest/gtest.h>

namespace lavaencrypt {

TEST(EncryptorDeriveKey, EmptyEntropyThrows) {
    const Encryptor encryptor;
    EXPECT_THROW(encryptor.deriveKey({}), std::invalid_argument);
}

TEST(EncryptorDeriveKey, NotYetImplementedThrows) {
    const Encryptor encryptor;
    const std::vector<uint8_t> entropy(32, 0xAB);
    EXPECT_THROW(encryptor.deriveKey(entropy), std::exception);
}

TEST(EncryptorEncrypt, NotYetImplementedThrows) {
    const Encryptor encryptor;
    const std::vector<uint8_t> plaintext = {0x01, 0x02, 0x03};
    const std::vector<uint8_t> entropy(32, 0xCD);
    EXPECT_THROW(encryptor.encrypt(plaintext, entropy), std::exception);
}

TEST(EncryptorDecrypt, NotYetImplementedThrows) {
    const Encryptor encryptor;
    const std::vector<uint8_t> ciphertext = {0x01, 0x02};
    const std::vector<uint8_t> key(32, 0x00);
    EXPECT_THROW(encryptor.decrypt(ciphertext, key), std::exception);
}

} // namespace lavaencrypt
