#include "lavaencrypt/utils.h"

#include <fstream>
#include <iomanip>
#include <sstream>
#include <stdexcept>

namespace lavaencrypt::utils {

std::vector<uint8_t> readFile(const std::string& filepath) {
    std::ifstream file(filepath, std::ios::binary | std::ios::ate);
    if (!file) {
        throw std::runtime_error("Cannot open file for reading: " + filepath);
    }
    const auto size = file.tellg();
    file.seekg(0);
    std::vector<uint8_t> data(static_cast<std::size_t>(size));
    file.read(reinterpret_cast<char*>(data.data()), size);
    return data;
}

void writeFile(const std::string& filepath, const std::vector<uint8_t>& data) {
    std::ofstream file(filepath, std::ios::binary);
    if (!file) {
        throw std::runtime_error("Cannot open file for writing: " + filepath);
    }
    file.write(reinterpret_cast<const char*>(data.data()),
               static_cast<std::streamsize>(data.size()));
}

std::string toHexString(const std::vector<uint8_t>& bytes) {
    std::ostringstream oss;
    oss << std::hex << std::setfill('0');
    for (const auto byte : bytes) {
        oss << std::setw(2) << static_cast<unsigned>(byte);
    }
    return oss.str();
}

std::vector<uint8_t> fromHexString(const std::string& hex) {
    if (hex.size() % 2 != 0) {
        throw std::invalid_argument("Hex string must have even length");
    }
    std::vector<uint8_t> bytes;
    bytes.reserve(hex.size() / 2);
    for (std::size_t i = 0; i < hex.size(); i += 2) {
        const std::string byteStr = hex.substr(i, 2);
        // Validate characters
        for (const char c : byteStr) {
            if (!std::isxdigit(static_cast<unsigned char>(c))) {
                throw std::invalid_argument(
                    std::string("Invalid hex character: ") + c);
            }
        }
        bytes.push_back(static_cast<uint8_t>(std::stoul(byteStr, nullptr, 16)));
    }
    return bytes;
}

} // namespace lavaencrypt::utils
