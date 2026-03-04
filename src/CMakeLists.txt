# ---------------------------------------------------------------------------
# Core library (object library to avoid linking the library twice)
# ---------------------------------------------------------------------------
add_library(lavaencrypt_lib STATIC
    image_processor.cpp
    entropy_extractor.cpp
    encryptor.cpp
    utils.cpp
)

target_include_directories(lavaencrypt_lib
    PUBLIC
        ${PROJECT_SOURCE_DIR}/include
)

target_compile_options(lavaencrypt_lib PRIVATE
    $<$<CXX_COMPILER_ID:GNU,Clang,AppleClang>:-Wall -Wextra -Wpedantic>
    $<$<CXX_COMPILER_ID:MSVC>:/W4>
)

# ---------------------------------------------------------------------------
# CLI executable
# ---------------------------------------------------------------------------
add_executable(lavaencrypt main.cpp)

target_link_libraries(lavaencrypt PRIVATE lavaencrypt_lib)

install(TARGETS lavaencrypt RUNTIME DESTINATION bin)
