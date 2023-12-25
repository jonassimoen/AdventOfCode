# Create a library with the input functions
add_library(util input.cpp input.h grid.h)
target_include_directories(util PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

if (CMAKE_VERSION VERSION_GREATER 3.12)
  set_property(TARGET util PROPERTY CXX_STANDARD 20)
endif()
