# Copyright 2019 Tricot Inc.
# Use of this source code is governed by the license in the LICENSE file.

cmake_minimum_required(VERSION 3.10.0)

cmake_policy(SET CMP0048 NEW)

enable_testing()

function(_bazel2cmake_cc_config name scope)
    target_compile_features("${name}" "${scope}" cxx_std_11)
    target_include_directories("${name}" "${scope}" "${PROJECT_SOURCE_DIR}")
endfunction()

function(bazel2cmake_cc_library name)
    if(DEFINED BAZEL2CMAKE_SKIP_TARGET_REGEX AND name MATCHES "${BAZEL2CMAKE_SKIP_TARGET_REGEX}")
        return()
    endif()

    cmake_parse_arguments("arg" "" "" "SRCS;HDRS;DEPS" ${ARGN})
    if(arg_SRCS)
        set(scope "PUBLIC")
    else()
        # Header-only library (or it might not have any headers).
        set(scope "INTERFACE")
        set(arg_SRCS "INTERFACE")
    endif()

    # TODO(vtl): Do something with arg_HDRS?
    add_library("${name}" ${arg_SRCS})
    _bazel2cmake_cc_config("${name}" "${scope}")
    target_link_libraries("${name}" "${scope}" ${arg_DEPS})
endfunction() 

function(bazel2cmake_cc_binary name)
    if(DEFINED BAZEL2CMAKE_SKIP_TARGET_REGEX AND name MATCHES "${BAZEL2CMAKE_SKIP_TARGET_REGEX}")
        return()
    endif()

    cmake_parse_arguments("arg" "" "" "SRCS;DEPS" ${ARGN})

    add_executable("${name}" ${arg_SRCS})
    _bazel2cmake_cc_config("${name}" PRIVATE)
    target_link_libraries("${name}" ${arg_DEPS})
endfunction() 

function(bazel2cmake_cc_test name)
    if(DEFINED BAZEL2CMAKE_SKIP_TARGET_REGEX AND name MATCHES "${BAZEL2CMAKE_SKIP_TARGET_REGEX}")
        return()
    endif()

    bazel2cmake_cc_binary("${name}" ${ARGN})
    add_test(NAME "${name}" COMMAND "${name}")
endfunction()
