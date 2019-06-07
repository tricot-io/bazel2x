# Copyright 2019 Tricot Inc.
# Use of this source code is governed by the license in the LICENSE file.

cmake_minimum_required(VERSION 3.10.0)

cmake_policy(SET CMP0048 NEW)

enable_testing()

function(_bazel2cmake_skip_target out name)
    if(DEFINED BAZEL2CMAKE_SKIP_TARGET_REGEX AND name MATCHES "${BAZEL2CMAKE_SKIP_TARGET_REGEX}")
        set("${out}" TRUE PARENT_SCOPE)
    else()
        set("${out}" FALSE PARENT_SCOPE)
    endif()
endfunction()

function(_bazel2cmake_cc_config name scope)
    target_compile_features("${name}" "${scope}" cxx_std_11)
    target_include_directories("${name}" "${scope}" "${PROJECT_SOURCE_DIR}")
endfunction()

function(bazel2cmake_cc_library name)
    _bazel2cmake_skip_target(skip "${name}")
    if(skip)
        return()
    endif()

    set(srcs "")
    set(hdrs "")
    set(deps "")
    set(attr "NONE")
    foreach(arg IN LISTS ARGN)
        if(arg STREQUAL "SRCS")
            set(attr "SRCS")
        elseif(arg STREQUAL "HDRS")
            set(attr "HDRS")
        elseif(arg STREQUAL "DEPS")
            set(attr "DEPS")
        else()
            if(attr STREQUAL "NONE")
                message(FATAL_ERROR "${name}: unknown or missing attribute ${arg}")
            elseif(attr STREQUAL "SRCS")
                list(APPEND srcs "${arg}")
            elseif(attr STREQUAL "HDRS")
                list(APPEND hdrs "${arg}")
            elseif(attr STREQUAL "DEPS")
                list(APPEND deps "${arg}")
            endif()
        endif()
    endforeach()
    if(srcs STREQUAL "")
        # Header-only library (or it might not have any headers).
        set(scope "INTERFACE")
        set(srcs "INTERFACE")
    else()
        set(scope "PUBLIC")
    endif()
    # deps is allowed to be empty.

    add_library("${name}" ${srcs})
    _bazel2cmake_cc_config("${name}" "${scope}")
    target_link_libraries("${name}" "${scope}" ${deps})
endfunction() 

function(bazel2cmake_cc_binary name)
    _bazel2cmake_skip_target(skip "${name}")
    if(skip)
        return()
    endif()

    set(srcs "")
    set(deps "")
    set(attr "NONE")
    foreach(arg IN LISTS ARGN)
        if(arg STREQUAL "SRCS")
            set(attr "SRCS")
        elseif(arg STREQUAL "DEPS")
            set(attr "DEPS")
        else()
            if(attr STREQUAL "NONE")
                message(FATAL_ERROR "${name}: unknown or missing attribute ${arg}")
            elseif(attr STREQUAL "SRCS")
                list(APPEND srcs "${arg}")
            elseif(attr STREQUAL "DEPS")
                list(APPEND deps "${arg}")
            endif()
        endif()
    endforeach()

    add_executable("${name}" ${srcs})
    _bazel2cmake_cc_config("${name}" PRIVATE)
    target_link_libraries("${name}" ${deps})
endfunction() 

function(bazel2cmake_cc_test name)
    _bazel2cmake_skip_target(skip "${name}")
    if(skip)
        return()
    endif()

    bazel2cmake_cc_binary("${name}" ${ARGN})
    add_test(NAME "${name}" COMMAND "${name}")
endfunction()
