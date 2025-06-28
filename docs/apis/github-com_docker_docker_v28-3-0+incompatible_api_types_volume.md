# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/api/types/volume

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:26:32 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

This section is empty.

## Types

### AccessMode

AccessMode defines the access mode of a volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L86)  

```go
type AccessMode struct {
	// Scope defines the set of nodes this volume can be used on at one time.
	Scope Scope `json:",omitempty"`

	// Sharing defines the number and way that different tasks can use this
	// volume at one time.
	Sharing SharingMode `json:",omitempty"`

	// MountVolume defines options for using this volume as a Mount-type
	// volume.
	//
	// Either BlockVolume or MountVolume, but not both, must be present.
	MountVolume *TypeMount `json:",omitempty"`

	// BlockVolume defines options for using this volume as a Block-type
	// volume.
	//
	// Either BlockVolume or MountVolume, but not both, must be present.
	BlockVolume *TypeBlock `json:",omitempty"`
}
```

---

### Availability

Availability specifies the availability of the volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L69)  

```go
type Availability string
```

---

### CapacityRange

CapacityRange describes the minimum and maximum capacity a volume should be
created with

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L339)  

```go
type CapacityRange struct {
	// RequiredBytes specifies that a volume must be at least this big. The
	// value of 0 indicates an unspecified minimum.
	RequiredBytes int64

	// LimitBytes specifies that a volume must not be bigger than this. The
	// value of 0 indicates an unspecified maximum
	LimitBytes int64
}
```

---

### ClusterVolume

ClusterVolume contains options and information specific to, and only present
on, Swarm CSI cluster volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L9)  

```go
type ClusterVolume struct {
	// ID is the Swarm ID of the volume. Because cluster volumes are Swarm
	// objects, they have an ID, unlike non-cluster volumes, which only have a
	// Name. This ID can be used to refer to the cluster volume.
	ID string

	// Meta is the swarm metadata about this volume.
	swarm.Meta

	// Spec is the cluster-specific options from which this volume is derived.
	Spec ClusterVolumeSpec

	// PublishStatus contains the status of the volume as it pertains to its
	// publishing on Nodes.
	PublishStatus []*PublishStatus `json:",omitempty"`

	// Info is information about the global status of the volume.
	Info *Info `json:",omitempty"`
}
```

---

### ClusterVolumeSpec

ClusterVolumeSpec contains the spec used to create this volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L30)  

```go
type ClusterVolumeSpec struct {
	// Group defines the volume group of this volume. Volumes belonging to the
	// same group can be referred to by group name when creating Services.
	// Referring to a volume by group instructs swarm to treat volumes in that
	// group interchangeably for the purpose of scheduling. Volumes with an
	// empty string for a group technically all belong to the same, emptystring
	// group.
	Group string `json:",omitempty"`

	// AccessMode defines how the volume is used by tasks.
	AccessMode *AccessMode `json:",omitempty"`

	// AccessibilityRequirements specifies where in the cluster a volume must
	// be accessible from.
	//
	// This field must be empty if the plugin does not support
	// VOLUME_ACCESSIBILITY_CONSTRAINTS capabilities. If it is present but the
	// plugin does not support it, volume will not be created.
	//
	// If AccessibilityRequirements is empty, but the plugin does support
	// VOLUME_ACCESSIBILITY_CONSTRAINTS, then Swarmkit will assume the entire
	// cluster is a valid target for the volume.
	AccessibilityRequirements *TopologyRequirement `json:",omitempty"`

	// CapacityRange defines the desired capacity that the volume should be
	// created with. If nil, the plugin will decide the capacity.
	CapacityRange *CapacityRange `json:",omitempty"`

	// Secrets defines Swarm Secrets that are passed to the CSI storage plugin
	// when operating on this volume.
	Secrets []Secret `json:",omitempty"`

	// Availability is the Volume's desired availability. Analogous to Node
	// Availability, this allows the user to take volumes offline in order to
	// update or delete them.
	Availability Availability `json:",omitempty"`
}
```

---

### CreateOptions

CreateOptions VolumeConfig

Volume configuration
swagger:model CreateOptions

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/create_options.go#L10)  

```go
type CreateOptions struct {

	// cluster volume spec
	ClusterVolumeSpec *ClusterVolumeSpec `json:"ClusterVolumeSpec,omitempty"`

	// Name of the volume driver to use.
	Driver string `json:"Driver,omitempty"`

	// A mapping of driver options and values. These options are
	// passed directly to the driver and are driver specific.
	//
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// The new volume's name. If not specified, Docker generates a name.
	//
	Name string `json:"Name,omitempty"`
}
```

---

### DiskUsage

DiskUsage contains disk usage for volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/disk_usage.go#L4)  

```go
type DiskUsage struct {
	TotalSize   int64
	Reclaimable int64
	Items       []*Volume
}
```

---

### Info

Info contains information about the Volume as a whole as provided by
the CSI storage plugin.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L402)  

```go
type Info struct {
	// CapacityBytes is the capacity of the volume in bytes. A value of 0
	// indicates that the capacity is unknown.
	CapacityBytes int64 `json:",omitempty"`

	// VolumeContext is the context originating from the CSI storage plugin
	// when the Volume is created.
	VolumeContext map[string]string `json:",omitempty"`

	// VolumeID is the ID of the Volume as seen by the CSI storage plugin. This
	// is distinct from the Volume's Swarm ID, which is the ID used by all of
	// the Docker Engine to refer to the Volume. If this field is blank, then
	// the Volume has not been successfully created yet.
	VolumeID string `json:",omitempty"`

	// AccessibleTopology is the topology this volume is actually accessible
	// from.
	AccessibleTopology []Topology `json:",omitempty"`
}
```

---

### ListOptions

ListOptions holds parameters to list volumes.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/options.go#L6)  

```go
type ListOptions struct {
	Filters filters.Args
}
```

---

### ListResponse

ListResponse VolumeListResponse

Volume list response
swagger:model ListResponse

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/list_response.go#L10)  

```go
type ListResponse struct {

	// List of volumes
	Volumes []*Volume `json:"Volumes"`

	// Warnings that occurred when fetching the list of volumes.
	//
	Warnings []string `json:"Warnings"`
}
```

---

### PruneReport

PruneReport contains the response for Engine API:
POST "/volumes/prune"

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/options.go#L12)  

```go
type PruneReport struct {
	VolumesDeleted []string
	SpaceReclaimed uint64
}
```

---

### PublishState

PublishState represents the state of a Volume as it pertains to its
use on a particular Node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L364)  

```go
type PublishState string
```

---

### PublishStatus

PublishStatus represents the status of the volume as published to an
individual node

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L388)  

```go
type PublishStatus struct {
	// NodeID is the ID of the swarm node this Volume is published to.
	NodeID string `json:",omitempty"`

	// State is the publish state of the volume.
	State PublishState `json:",omitempty"`

	// PublishContext is the PublishContext returned by the CSI plugin when
	// a volume is published.
	PublishContext map[string]string `json:",omitempty"`
}
```

---

### Scope

Scope defines the Scope of a Cluster Volume. This is how many nodes a
Volume can be accessed simultaneously on.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L109)  

```go
type Scope string
```

---

### Secret

Secret represents a Swarm Secret value that must be passed to the CSI
storage plugin when operating on this Volume. It represents one key-value
pair of possibly many.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L352)  

```go
type Secret struct {
	// Key is the name of the key of the key-value pair passed to the plugin.
	Key string

	// Secret is the swarm Secret object from which to read data. This can be a
	// Secret name or ID. The Secret data is retrieved by Swarm and used as the
	// value of the key-value pair passed to the plugin.
	Secret string
}
```

---

### SharingMode

SharingMode defines the Sharing of a Cluster Volume. This is how Tasks using a
Volume at the same time can use it.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L123)  

```go
type SharingMode string
```

---

### Topology

Topology is a map of topological domains to topological segments.

This description is taken verbatim from the CSI Spec:

A topological domain is a sub-division of a cluster, like "region",
"zone", "rack", etc.
A topological segment is a specific instance of a topological domain,
like "zone3", "rack3", etc.
For example {"com.company/zone": "Z1", "com.company/rack": "R3"}
Valid keys have two segments: an OPTIONAL prefix and name, separated
by a slash (/), for example: "com.company.example/zone".
The key name segment is REQUIRED. The prefix is OPTIONAL.
The key name MUST be 63 characters or less, begin and end with an
alphanumeric character ([a-z0-9A-Z]), and contain only dashes (-),
underscores (_), dots (.), or alphanumerics in between, for example
"zone".
The key prefix MUST be 63 characters or less, begin and end with a
lower-case alphanumeric character ([a-z0-9]), contain only
dashes (-), dots (.), or lower-case alphanumerics in between, and
follow domain name notation format
(https://tools.ietf.org/html/rfc1035#section-2.3.1).
The key prefix SHOULD include the plugin's host company name and/or
the plugin name, to minimize the possibility of collisions with keys
from other plugins.
If a key prefix is specified, it MUST be identical across all
topology keys returned by the SP (across all RPCs).
Keys MUST be case-insensitive. Meaning the keys "Zone" and "zone"
MUST not both exist.
Each value (topological segment) MUST contain 1 or more strings.
Each string MUST be 63 characters or less and begin and end with an
alphanumeric character with '-', '_', '.', or alphanumerics in
between.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L333)  

```go
type Topology struct {
	Segments map[string]string `json:",omitempty"`
}
```

---

### TopologyRequirement

TopologyRequirement expresses the user's requirements for a volume's
accessible topology.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L160)  

```go
type TopologyRequirement struct {
	// Requisite specifies a list of Topologies, at least one of which the
	// volume must be accessible from.
	//
	// Taken verbatim from the CSI Spec:
	//
	// Specifies the list of topologies the provisioned volume MUST be
	// accessible from.
	// This field is OPTIONAL. If TopologyRequirement is specified either
	// requisite or preferred or both MUST be specified.
	//
	// If requisite is specified, the provisioned volume MUST be
	// accessible from at least one of the requisite topologies.
	//
	// Given
	//   x = number of topologies provisioned volume is accessible from
	//   n = number of requisite topologies
	// The CO MUST ensure n >= 1. The SP MUST ensure x >= 1
	// If x==n, then the SP MUST make the provisioned volume available to
	// all topologies from the list of requisite topologies. If it is
	// unable to do so, the SP MUST fail the CreateVolume call.
	// For example, if a volume should be accessible from a single zone,
	// and requisite =
	//   {"region": "R1", "zone": "Z2"}
	// then the provisioned volume MUST be accessible from the "region"
	// "R1" and the "zone" "Z2".
	// Similarly, if a volume should be accessible from two zones, and
	// requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"}
	// then the provisioned volume MUST be accessible from the "region"
	// "R1" and both "zone" "Z2" and "zone" "Z3".
	//
	// If x<n, then the SP SHALL choose x unique topologies from the list
	// of requisite topologies. If it is unable to do so, the SP MUST fail
	// the CreateVolume call.
	// For example, if a volume should be accessible from a single zone,
	// and requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"}
	// then the SP may choose to make the provisioned volume available in
	// either the "zone" "Z2" or the "zone" "Z3" in the "region" "R1".
	// Similarly, if a volume should be accessible from two zones, and
	// requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"},
	//   {"region": "R1", "zone": "Z4"}
	// then the provisioned volume MUST be accessible from any combination
	// of two unique topologies: e.g. "R1/Z2" and "R1/Z3", or "R1/Z2" and
	//  "R1/Z4", or "R1/Z3" and "R1/Z4".
	//
	// If x>n, then the SP MUST make the provisioned volume available from
	// all topologies from the list of requisite topologies and MAY choose
	// the remaining x-n unique topologies from the list of all possible
	// topologies. If it is unable to do so, the SP MUST fail the
	// CreateVolume call.
	// For example, if a volume should be accessible from two zones, and
	// requisite =
	//   {"region": "R1", "zone": "Z2"}
	// then the provisioned volume MUST be accessible from the "region"
	// "R1" and the "zone" "Z2" and the SP may select the second zone
	// independently, e.g. "R1/Z4".
	Requisite []Topology `json:",omitempty"`

	// Preferred is a list of Topologies that the volume should attempt to be
	// provisioned in.
	//
	// Taken from the CSI spec:
	//
	// Specifies the list of topologies the CO would prefer the volume to
	// be provisioned in.
	//
	// This field is OPTIONAL. If TopologyRequirement is specified either
	// requisite or preferred or both MUST be specified.
	//
	// An SP MUST attempt to make the provisioned volume available using
	// the preferred topologies in order from first to last.
	//
	// If requisite is specified, all topologies in preferred list MUST
	// also be present in the list of requisite topologies.
	//
	// If the SP is unable to make the provisioned volume available
	// from any of the preferred topologies, the SP MAY choose a topology
	// from the list of requisite topologies.
	// If the list of requisite topologies is not specified, then the SP
	// MAY choose from the list of all possible topologies.
	// If the list of requisite topologies is specified and the SP is
	// unable to make the provisioned volume available from any of the
	// requisite topologies it MUST fail the CreateVolume call.
	//
	// Example 1:
	// Given a volume should be accessible from a single zone, and
	// requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"}
	// preferred =
	//   {"region": "R1", "zone": "Z3"}
	// then the SP SHOULD first attempt to make the provisioned volume
	// available from "zone" "Z3" in the "region" "R1" and fall back to
	// "zone" "Z2" in the "region" "R1" if that is not possible.
	//
	// Example 2:
	// Given a volume should be accessible from a single zone, and
	// requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"},
	//   {"region": "R1", "zone": "Z4"},
	//   {"region": "R1", "zone": "Z5"}
	// preferred =
	//   {"region": "R1", "zone": "Z4"},
	//   {"region": "R1", "zone": "Z2"}
	// then the SP SHOULD first attempt to make the provisioned volume
	// accessible from "zone" "Z4" in the "region" "R1" and fall back to
	// "zone" "Z2" in the "region" "R1" if that is not possible. If that
	// is not possible, the SP may choose between either the "zone"
	// "Z3" or "Z5" in the "region" "R1".
	//
	// Example 3:
	// Given a volume should be accessible from TWO zones (because an
	// opaque parameter in CreateVolumeRequest, for example, specifies
	// the volume is accessible from two zones, aka synchronously
	// replicated), and
	// requisite =
	//   {"region": "R1", "zone": "Z2"},
	//   {"region": "R1", "zone": "Z3"},
	//   {"region": "R1", "zone": "Z4"},
	//   {"region": "R1", "zone": "Z5"}
	// preferred =
	//   {"region": "R1", "zone": "Z5"},
	//   {"region": "R1", "zone": "Z3"}
	// then the SP SHOULD first attempt to make the provisioned volume
	// accessible from the combination of the two "zones" "Z5" and "Z3" in
	// the "region" "R1". If that's not possible, it should fall back to
	// a combination of "Z5" and other possibilities from the list of
	// requisite. If that's not possible, it should fall back  to a
	// combination of "Z3" and other possibilities from the list of
	// requisite. If that's not possible, it should fall back  to a
	// combination of other possibilities from the list of requisite.
	Preferred []Topology `json:",omitempty"`
}
```

---

### TypeBlock

TypeBlock defines options for using a volume as a block-type volume.

Intentionally empty.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L146)  

```go
type TypeBlock struct{}
```

---

### TypeMount

TypeMount contains options for using a volume as a Mount-type
volume.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/cluster_volume.go#L150)  

```go
type TypeMount struct {
	// FsType specifies the filesystem type for the mount volume. Optional.
	FsType string `json:",omitempty"`

	// MountFlags defines flags to pass when mounting the volume. Optional.
	MountFlags []string `json:",omitempty"`
}
```

---

### UpdateOptions

UpdateOptions is configuration to update a Volume with.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/volume_update.go#L4)  

```go
type UpdateOptions struct {
	// Spec is the ClusterVolumeSpec to update the volume to.
	Spec *ClusterVolumeSpec `json:"Spec,omitempty"`
}
```

---

### UsageData

UsageData Usage details about the volume. This information is used by the
`GET /system/df` endpoint, and omitted in other endpoints.

swagger:model UsageData

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/volume.go#L60)  

```go
type UsageData struct {

	// The number of containers referencing this volume. This field
	// is set to `-1` if the reference-count is not available.
	//
	// Required: true
	RefCount int64 `json:"RefCount"`

	// Amount of disk space used by the volume (in bytes). This information
	// is only available for volumes created with the `"local"` volume
	// driver. For volumes created with other volume drivers, this field
	// is set to `-1` ("not available")
	//
	// Required: true
	Size int64 `json:"Size"`
}
```

---

### Volume

Volume volume
swagger:model Volume

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/api/types/volume/volume.go#L8)  

```go
type Volume struct {

	// cluster volume
	ClusterVolume *ClusterVolume `json:"ClusterVolume,omitempty"`

	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// Name of the volume driver used by the volume.
	// Required: true
	Driver string `json:"Driver"`

	// User-defined key/value metadata.
	// Required: true
	Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	// Required: true
	Mountpoint string `json:"Mountpoint"`

	// Name of the volume.
	// Required: true
	Name string `json:"Name"`

	// The driver specific options used when creating the volume.
	//
	// Required: true
	Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide,
	// or `local` for machine level.
	//
	// Required: true
	Scope string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	//
	Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	UsageData *UsageData `json:"UsageData,omitempty"`
}
```

---

