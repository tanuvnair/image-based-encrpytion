#pragma once

#include <cstdint>
#include <string>
#include <vector>

namespace lavaencrypt {

/// Configuration for the encryption operation.
struct EncryptorConfig {
    std::size_t keyLengthBytes = 32; ///< Desired symmetric key length (default: 256-bit)
};

/// Uses the entropy bytes produced by EntropyExtractor to derive a symmetric
/// key and encrypt an arbitrary payload.
///
/// The exact cipher algorithm (AES-GCM, ChaCha20-Poly1305, …) is intentionally
/// left as an implementation detail; the interface is algorithm-agnostic.
class Encryptor {
public:
    explicit Encryptor(const EncryptorConfig& config = {});
    ~Encryptor() = default;

    /// Derive a symmetric key from @p entropyBytes.
    /// @throws std::invalid_argument if @p entropyBytes is empty.
    std::vector<uint8_t> deriveKey(const std::vector<uint8_t>& entropyBytes) const;

    /// Encrypt @p plaintext using a key derived from @p entropyBytes.
    /// @return Ciphertext (including any authentication tag / nonce prefix).
    std::vector<uint8_t> encrypt(
        const std::vector<uint8_t>& plaintext,
        const std::vector<uint8_t>& entropyBytes) const;

    /// Decrypt @p ciphertext using @p key.
    /// @return Recovered plaintext.
    /// @throws std::runtime_error if decryption or authentication fails.
    std::vector<uint8_t> decrypt(
        const std::vector<uint8_t>& ciphertext,
        const std::vector<uint8_t>& key) const;

private:
    EncryptorConfig m_config;
};

} // namespace lavaencrypt
