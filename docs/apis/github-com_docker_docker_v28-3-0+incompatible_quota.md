# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/quota

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:37:01 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/errors.go#L8)

```go
var ErrQuotaNotSupported = errQuotaNotSupported{}
```

## Functions

### CanTestQuota

CanTestQuota - checks if xfs prjquota can be tested
returns a reason if not

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/testhelpers.go#L15)  

```go
func CanTestQuota() (string, bool)
```

---

### PrepareQuotaTestImage

PrepareQuotaTestImage - prepares an xfs prjquota test image
returns the path of the image on success

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/testhelpers.go#L28)  

```go
func PrepareQuotaTestImage(t *testing.T) (string, error)
```

---

### WrapMountTest

WrapMountTest - wraps a test function such that it has easy access to a mountPoint and testDir
with guaranteed prjquota or guaranteed no prjquota support.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/testhelpers.go#L75)  

```go
func WrapMountTest(imageFileName string, enableQuota bool, testFunc func(t *testing.T, mountPoint, backingFsDev, testDir string)) func(*testing.T)
```

---

### WrapQuotaTest

WrapQuotaTest - wraps a test function such that is has easy and guaranteed access to a quota Control
instance with a quota test dir under its control.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/testhelpers.go#L119)  

```go
func WrapQuotaTest(testFunc func(t *testing.T, ctrl *Control, mountPoint, testDir, testSubDir string)) func(t *testing.T, mountPoint, backingFsDev, testDir string)
```

---

## Types

### Control

Control - Context to be used by storage driver (e.g. overlay)
who wants to apply project quotas to container dirs

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/types.go#L12)  

```go
type Control struct {
	sync.RWMutex // protect nextProjectID and quotas map
	// contains filtered or unexported fields
}
```

#### Functions

##### NewControl

NewControl - initialize project quota support.
Test to make sure that quota can be set on a test dir and find
the first project id to be used for the next container create.

Returns nil (and error) if project quota is not supported.

First get the project id of the home directory.
This test will fail if the backing fs is not xfs.

xfs_quota tool can be used to assign a project id to the driver home directory, e.g.:

In that case, the home directory project id will be used as a "start offset"
and all containers will be assigned larger project ids (e.g. >= 1000).
This is a way to prevent xfs_quota management from conflicting with docker.

Then try to create a test directory with the next project id and set a quota
on it. If that works, continue to scan existing containers to map allocated
project ids.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/projectquota.go#L120)  

```go
func NewControl(basePath string) (*Control, error)
```

#### Methods

##### Control.GetQuota

GetQuota - get the quota limits of a directory that was configured with SetQuota

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/projectquota.go#L251)  

```go
func (q *Control) GetQuota(targetPath string, quota *Quota) error
```

##### Control.SetQuota

SetQuota - assign a unique project id to directory and set the quota limits
for that project id

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/projectquota.go#L192)  

```go
func (q *Control) SetQuota(targetPath string, quota Quota) error
```

---

### Quota

Quota limit params - currently we only control blocks hard limit

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/quota/types.go#L6)  

```go
type Quota struct {
	Size uint64
}
```

---

