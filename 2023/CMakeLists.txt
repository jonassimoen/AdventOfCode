# CMakeList.txt : CMake project for 2023, include source and define
# project specific logic here.
cmake_minimum_required(VERSION 3.10)
project ("AdventOfCode2023")

# Add executable that uses the util library
add_subdirectory(util)
add_executable(2023 2023.cpp)
target_link_libraries(2023 util)

add_definitions(-DINPUT_FOLDER="${CMAKE_CURRENT_SOURCE_DIR}/input")
add_definitions(-DASSETS_FOLDER="${CMAKE_CURRENT_SOURCE_DIR}/assets")

if (CMAKE_VERSION VERSION_GREATER 3.12)
  set_property(TARGET 2023 PROPERTY CXX_STANDARD 23)
endif()

