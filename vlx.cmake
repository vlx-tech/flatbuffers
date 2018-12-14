cmake_minimum_required(VERSION 3.13)
include_guard(GLOBAL)

option(FLATBUFFERS_BUILD_TESTS OFF)
option(FLATBUFFERS_INSTALL OFF)
option(FLATBUFFERS_BUILD_FLATC OFF)
option(FLATBUFFERS_BUILD_FLATHASH OFF)

set(FLATBUFFERS_FLATC_EXECUTABLE "${ROOT}/tools/flatbuffers/bin/${CMAKE_HOST_SYSTEM_NAME}/flatc" CACHE INTERNAL "")

add_subdirectory(
    ${CMAKE_CURRENT_LIST_DIR}
    "external/flatbuffers"
)
