#include "lavaencrypt/encryptor.h"

#include <stdexcept>

namespace lavaencrypt {

Encryptor::Encryptor(const EncryptorConfig& config) : m_config(config) {}

std::vector<uint8_t> Encryptor::deriveKey(
    const std::vector<uint8_t>& entropyBytes) const {
    if (entropyBytes.empty()) {
        throw std::invalid_argument("entropyBytes must not be empty");
    }
    // TODO: apply HKDF or similar KDF to derive a fixed-length symmetric key.
    throw std::runtime_error("Encryptor::deriveKey not yet implemented");
}

std::vector<uint8_t> Encryptor::encrypt(
    const std::vector<uint8_t>& /*plaintext*/,
    const std::vector<uint8_t>& entropyBytes) const {
    // TODO: derive key then encrypt with AES-GCM or ChaCha20-Poly1305.
    (void)deriveKey(entropyBytes); // validates input early
    throw std::runtime_error("Encryptor::encrypt not yet implemented");
}

std::vector<uint8_t> Encryptor::decrypt(
    const std::vector<uint8_t>& /*ciphertext*/,
    const std::vector<uint8_t>& /*key*/) const {
    // TODO: decrypt and authenticate ciphertext with the provided key.
    throw std::runtime_error("Encryptor::decrypt not yet implemented");
}

} // namespace lavaencrypt
