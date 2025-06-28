# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/cli/winresources/dockerd

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:11 UTC

## Overview

Package winresources is used to embed Windows resources into dockerd.exe.

These resources are used to provide:
* Version information
* An icon
* A Windows manifest declaring Windows version support
* Events message table

The resource object files are generated when building with go-winres
in hack/make/.go-autogen and are located in cli/winresources.
This occurs automatically when you cross build against Windows OS.


