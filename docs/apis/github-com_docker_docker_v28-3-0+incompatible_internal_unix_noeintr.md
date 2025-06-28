# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/internal/unix_noeintr

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:32:09 UTC

## Overview

Package unix_noeintr provides wrappers for unix syscalls that retry on EINTR.

TODO: Consider moving (for example to moby/sys) and making the wrappers auto-generated.


## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Close

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L47)  

```go
func Close(fd int) (err error)
```

---

### EpollCreate

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/epoll_linux.go#L9)  

```go
func EpollCreate() (int, error)
```

---

### EpollCtl

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/epoll_linux.go#L19)  

```go
func EpollCtl(epFd int, op int, fd int, event *unix.EpollEvent) error
```

---

### EpollWait

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/epoll_linux.go#L29)  

```go
func EpollWait(epFd int, events []unix.EpollEvent, msec int) (int, error)
```

---

### Fstat

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L71)  

```go
func Fstat(fd int, stat *unix.Stat_t) (err error)
```

---

### Fstatat

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L79)  

```go
func Fstatat(fd int, path string, stat *unix.Stat_t, flags int) (err error)
```

---

### Mount

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L23)  

```go
func Mount(source string, target string, fstype string, flags uintptr, data string) (err error)
```

---

### Open

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L39)  

```go
func Open(path string, mode int, perm uint32) (fd int, err error)
```

---

### Openat

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L55)  

```go
func Openat(dirfd int, path string, mode int, perms uint32) (fd int, err error)
```

---

### Openat2

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L63)  

```go
func Openat2(dirfd int, path string, how *unix.OpenHow) (fd int, err error)
```

---

### Retry

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L14)  

```go
func Retry(f func() error)
```

---

### Unmount

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/internal/unix_noeintr/fs_unix.go#L31)  

```go
func Unmount(target string, flags int) (err error)
```

---

## Types

This section is empty.

