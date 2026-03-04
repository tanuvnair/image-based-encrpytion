cmake_minimum_required(VERSION 3.20)

project(
    lavaencrypt
    VERSION 0.1.0
    DESCRIPTION "Cloudflare lava lamp image-based entropy and encryption"
    LANGUAGES CXX
)

# ---------------------------------------------------------------------------
# Global settings
# ---------------------------------------------------------------------------
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Produce compile_commands.json for IDE tooling
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# ---------------------------------------------------------------------------
# Options
# ---------------------------------------------------------------------------
option(LAVAENCRYPT_BUILD_TESTS "Build unit tests" ON)
option(LAVAENCRYPT_BUILD_DOCS  "Build documentation (requires Doxygen)" OFF)

# ---------------------------------------------------------------------------
# Third-party dependencies (fetched automatically if not found locally)
# ---------------------------------------------------------------------------
include(FetchContent)

if(LAVAENCRYPT_BUILD_TESTS)
    FetchContent_Declare(
        googletest
        GIT_REPOSITORY https://github.com/google/googletest.git
        GIT_TAG        v1.14.0
    )
    # Prevent GTest from overriding our compiler/linker options
    set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
    FetchContent_MakeAvailable(googletest)
endif()

# ---------------------------------------------------------------------------
# Sub-directories
# ---------------------------------------------------------------------------
add_subdirectory(src)

if(LAVAENCRYPT_BUILD_TESTS)
    enable_testing()
    add_subdirectory(tests)
endif()
