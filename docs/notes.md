# bazel2cmake notes

## Facts and problems

*   There are 3 types of Bazel files: *BUILD* (*BUILD.bazel*) files, *.bzl*
    files, and *WORKSPACE* files.
*   To make sense of a *BUILD* file, you have to at least *find* the *WORKSPACE*
    file; this determines the meaning of "absolute" paths (labels).
*   It's tempting to entirely ignore other (external) workspaces, and indeed
    there are good reasons for doing so.
    *   In translating Bazel to CMake, you probably don't want to automatically
        cross workspace boundaries.
    *   Even if you actually depend on something from another workspace, it's
        quite possible that that workspace has a more natural, "native" CMake
        build.
    *   So you'd at least want to *support* mapping external dependencies to a
        CMake dependency of some other name.
*   All that said, it's hard to simply ignore external workspaces. E.g., they
    may provide macros that are used in your own *BUILD* files, in which case
    you at least need to be able to load *.bzl* files from them.
    *   Being able to fetch dependencies poses a significant problem, but
        luckily Bazel provides the `bazel fetch` and `bazel sync` commands.
*   *BUILD* files, *.bzl* files, and *WORKSPACE* files differ in what is
    permitted.
    *   Most glaringly, *BUILD* and *WORKSPACE* files use a restricted subset of
        Starlark (e.g., `**kwargs` is not allowed).

## Decisions

*   Don't read *WORKSPACE* files.
*   Assume that other workspaces are available under *bazel-\<project
    name\>/external* (where *\<project name\>* seems to be the name of the
    directory containing the *WORKSPACE* file).
*   This is enough to load *.bzl* files from other workspaces, as far as I can
    tell. (Presumably, any name mappings are *not* applied in this case, or so I
    hope. Or, rather, name mappings are only applied when a *BUILD* file is
    executed, or perhaps even later. It cannot be done at *.bzl* load time,
    because the interpreter has no way of knowing what strings are labels.)
