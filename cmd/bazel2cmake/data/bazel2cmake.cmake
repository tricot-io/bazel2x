# Copyright 2019 Tricot Inc.
# Use of this source code is governed by the license in the LICENSE file.

cmake_minimum_required(VERSION 3.10.0)

# NOTE: This is not just an include guard for this file, but for all bazel2cmake.cmake files
# (including in subprojects).
if(NOT _BAZEL2CMAKE_INITIALIZED)
    set(_BAZEL2CMAKE_INITIALIZED TRUE)
    include(cmake/bazel2cmake_defs.cmake)
    include(cmake/bazel2cmake_project.cmake OPTIONAL)
endif()
