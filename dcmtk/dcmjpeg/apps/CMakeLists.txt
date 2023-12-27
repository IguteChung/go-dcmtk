# GO-DCMTK dcmjpeg Project

# declare project
PROJECT(GO_DCMTK_DCMJPEG)

# minimum CMake version required
CMAKE_MINIMUM_REQUIRED(VERSION 2.8.3)

# include DCMTK's CMake configuration directly
INCLUDE(/usr/local/lib/cmake/dcmtk/DCMTKConfig.cmake)

# declare include directories
INCLUDE_DIRECTORIES(${DCMTK_INCLUDE_DIRS})

# declare executable and link required libraries
file(GLOB SOURCE_FILES "*.cc")
ADD_LIBRARY(godcmjpeg SHARED ${SOURCE_FILES})
TARGET_LINK_LIBRARIES(godcmjpeg ${DCMTK_LIBRARIES})

# Set the output directory for the shared library
set_target_properties(godcmjpeg PROPERTIES LIBRARY_OUTPUT_DIRECTORY "../dist")
