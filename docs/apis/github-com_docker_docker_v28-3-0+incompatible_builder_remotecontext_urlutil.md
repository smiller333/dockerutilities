# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/builder/remotecontext/urlutil

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:27:09 UTC

## Overview

Package urlutil provides helper function to check if a given build-context
location should be considered a URL or a remote Git repository.

This package is specifically written for use with docker build contexts, and
should not be used as a general-purpose utility.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### IsGitURL

IsGitURL returns true if the provided str is a remote git repository "URL".

This function only performs a rudimentary check (no validation is performed
to ensure the URL is well-formed), and is written specifically for use with
docker build, with some logic for backward compatibility with older versions
of docker: do not use this function as a general-purpose utility.

The following patterns are considered to be a Git URL:

The github.com/ prefix is a special case used to treat context-paths
starting with "github.com/" as a git URL if the given path does not
exist locally. The "github.com/" prefix is kept for backward compatibility,
and is a legacy feature.

Going forward, no additional prefixes should be added, and users should
be encouraged to use explicit URLs (https://github.com/user/repo.git) instead.

Note that IsGitURL does not check if "github.com/" prefixes exist as a local
path. Code using this function should check if the path exists locally before
using it as a URL.

Git URLs accept context configuration in their fragment section, separated by
a colon (`:`). The first part represents the reference to check out, and can
be either a branch, a tag, or a remote reference. The second part represents
a subdirectory inside the repository to use as the build context.

For example,the following URL uses a directory named "docker" in the branch
"container" in the https://github.com/myorg/my-repo.git repository:

https://github.com/myorg/my-repo.git#container:docker

The following table represents all the valid suffixes with their build
contexts:

| Build Syntax Suffix            | Git reference used   | Build Context Used |
|--------------------------------|----------------------|--------------------|
| my-repo.git                    | refs/heads/master    | /                  |
| my-repo.git#mytag              | refs/tags/my-tag     | /                  |
| my-repo.git#mybranch           | refs/heads/my-branch | /                  |
| my-repo.git#pull/42/head       | refs/pull/42/head    | /                  |
| my-repo.git#:directory         | refs/heads/master    | /directory         |
| my-repo.git#master:directory   | refs/heads/master    | /directory         |
| my-repo.git#mytag:directory    | refs/tags/my-tag     | /directory         |
| my-repo.git#mybranch:directory | refs/heads/my-branch | /directory         |

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/urlutil/urlutil.go#L77)  

```go
func IsGitURL(str string) bool
```

---

### IsURL

IsURL returns true if the provided str is an HTTP(S) URL by checking if it
has a http:// or https:// scheme. No validation is performed to verify if the
URL is well-formed.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/builder/remotecontext/urlutil/urlutil.go#L21)  

```go
func IsURL(str string) bool
```

---

## Types

This section is empty.

