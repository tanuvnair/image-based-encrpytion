#include "lavaencrypt/encryptor.h"
#include "lavaencrypt/entropy_extractor.h"
#include "lavaencrypt/image_processor.h"
#include "lavaencrypt/utils.h"

#include <iostream>
#include <stdexcept>

/// Entry point for the lavaencrypt CLI.
///
/// Usage:
///   lavaencrypt <input_image> <plaintext_file> <output_file>
///
/// The program:
///   1. Loads the input image and extracts entropy from its pixel data.
///   2. Uses that entropy to encrypt the contents of <plaintext_file>.
///   3. Writes the resulting ciphertext to <output_file>.
int main(int argc, char* argv[]) {
    if (argc != 4) {
        std::cerr << "Usage: " << argv[0]
                  << " <input_image> <plaintext_file> <output_file>\n";
        return 1;
    }

    const std::string imagePath     = argv[1];
    const std::string plaintextPath = argv[2];
    const std::string outputPath    = argv[3];

    try {
        // TODO: implement pipeline once modules are complete
        (void)imagePath;
        (void)plaintextPath;
        (void)outputPath;
        std::cout << "lavaencrypt: pipeline not yet implemented.\n";
    } catch (const std::exception& e) {
        std::cerr << "Error: " << e.what() << '\n';
        return 1;
    }

    return 0;
}
