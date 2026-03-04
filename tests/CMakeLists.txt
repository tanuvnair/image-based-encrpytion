include(GoogleTest)

add_executable(lavaencrypt_tests
    test_utils.cpp
    test_image_processor.cpp
    test_entropy_extractor.cpp
    test_encryptor.cpp
)

target_link_libraries(lavaencrypt_tests
    PRIVATE
        lavaencrypt_lib
        GTest::gtest_main
)

target_compile_options(lavaencrypt_tests PRIVATE
    $<$<CXX_COMPILER_ID:GNU,Clang,AppleClang>:-Wall -Wextra -Wpedantic>
    $<$<CXX_COMPILER_ID:MSVC>:/W4>
)

gtest_discover_tests(lavaencrypt_tests)
