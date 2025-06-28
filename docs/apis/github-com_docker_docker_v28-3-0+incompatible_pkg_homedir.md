# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/pkg/homedir

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:35:31 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Get

Get returns the home directory of the current user with the help of
environment variables depending on the target operating system.
Returned path should be used with "path/filepath" to form new paths.

On non-Windows platforms, it falls back to nss lookups, if the home
directory cannot be obtained from environment-variables.

If linking statically with cgo enabled against glibc, ensure the
osusergo build tag is used.

If needing to do nss lookups, do not disable cgo or set osusergo.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir.go#L20)  

```go
func Get() string
```

---

### GetConfigHome

GetConfigHome returns XDG_CONFIG_HOME.
GetConfigHome returns $HOME/.config and nil error if XDG_CONFIG_HOME is not set.
If HOME and XDG_CONFIG_HOME are not set, getpwent(3) is consulted to determine the users home directory.

See also https://standards.freedesktop.org/basedir-spec/latest/ar01s03.html

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir_linux.go#L86)  

```go
func GetConfigHome() (string, error)
```

---

### GetDataHome

GetDataHome returns XDG_DATA_HOME.
GetDataHome returns $HOME/.local/share and nil error if XDG_DATA_HOME is not set.
If HOME and XDG_DATA_HOME are not set, getpwent(3) is consulted to determine the users home directory.

See also https://standards.freedesktop.org/basedir-spec/latest/ar01s03.html

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir_linux.go#L70)  

```go
func GetDataHome() (string, error)
```

---

### GetLibHome

GetLibHome returns $HOME/.local/lib
If HOME is not set, getpwent(3) is consulted to determine the users home directory.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir_linux.go#L99)  

```go
func GetLibHome() (string, error)
```

---

### GetRuntimeDir

GetRuntimeDir returns XDG_RUNTIME_DIR.
XDG_RUNTIME_DIR is typically configured via pam_systemd.
GetRuntimeDir returns non-nil error if XDG_RUNTIME_DIR is not set.

See also https://standards.freedesktop.org/basedir-spec/latest/ar01s03.html

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir_linux.go#L15)  

```go
func GetRuntimeDir() (string, error)
```

---

### StickRuntimeDirContents

StickRuntimeDirContents sets the sticky bit on files that are under
XDG_RUNTIME_DIR, so that the files won't be periodically removed by the system.

StickyRuntimeDir returns slice of sticked files.
StickyRuntimeDir returns nil error if XDG_RUNTIME_DIR is not set.

See also https://standards.freedesktop.org/basedir-spec/latest/ar01s03.html

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/pkg/homedir/homedir_linux.go#L29)  

```go
func StickRuntimeDirContents(files []string) ([]string, error)
```

---

## Types

This section is empty.

