#pragma once

#include <cstdint>
#include <string>
#include <vector>

namespace lavaencrypt::utils {

/// Read the entire contents of a binary file into a byte vector.
/// @throws std::runtime_error if the file cannot be opened.
std::vector<uint8_t> readFile(const std::string& filepath);

/// Write @p data to @p filepath, creating or overwriting the file.
/// @throws std::runtime_error if the file cannot be written.
void writeFile(const std::string& filepath, const std::vector<uint8_t>& data);

/// Encode @p bytes as a lowercase hex string (e.g. for debug output).
std::string toHexString(const std::vector<uint8_t>& bytes);

/// Decode a hex string back to raw bytes.
/// @throws std::invalid_argument if @p hex has odd length or illegal characters.
std::vector<uint8_t> fromHexString(const std::string& hex);

} // namespace lavaencrypt::utils
