# tricot-io/bazel2x

## Overview

Bazel2x supports converting [Bazel](https://bazel.build/) BUILD files to other
formats. You should think of it as more of a toolkit for writing a converter
than a converter itself.

It consists of two parts. First, the "frontend" executes Bazel BUILD (and
WORKSPACE) files (loading .bzl files as necessary) and creates (rule) targets
(in memory). Second, a "backend" converter takes the in-memory targets and
produces output files.

Currently there is only one backend converter, which supports conversion to
CMake (CMakeLists.txt) files. This conversion is quirky and supports only our
use-cases (and is written by CMake novice); it is expected to evolve as the
requirements become better understood.

## TODOs

*   Tests.
*   Documentation.
*   More complete implementations of basic Bazel concepts including, e.g.:
    *   `glob()`
    *   configurations and `select()`
    *   `depset`s
    *   more (native) rules
    *   better understanding of different types of targets (e.g., rule targets
        vs file targets)
