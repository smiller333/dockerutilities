# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/daemon/cluster/convert

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 23:03:47 UTC

## Constants

This section is empty.

## Variables

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/service.go#L17)

```go
var (
	// ErrUnsupportedRuntime returns an error if the runtime is not supported by the daemon
	ErrUnsupportedRuntime = errors.New("unsupported runtime")
	// ErrMismatchedRuntime returns an error if the runtime does not match the provided spec
	ErrMismatchedRuntime = errors.New("mismatched Runtime and *Spec fields")
)
```

## Functions

### BasicNetworkCreateToGRPC

BasicNetworkCreateToGRPC converts a NetworkCreateRequest to a grpc NetworkSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/network.go#L188)  

```go
func BasicNetworkCreateToGRPC(create network.CreateRequest) swarmapi.NetworkSpec
```

---

### BasicNetworkFromGRPC

BasicNetworkFromGRPC converts a grpc Network to a NetworkResource.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/network.go#L140)  

```go
func BasicNetworkFromGRPC(n swarmapi.Network) network.Inspect
```

---

### ConfigFromGRPC

ConfigFromGRPC converts a grpc Config to a Config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/config.go#L10)  

```go
func ConfigFromGRPC(s *swarmapi.Config) swarmtypes.Config
```

---

### ConfigReferencesFromGRPC

ConfigReferencesFromGRPC converts a slice of grpc ConfigReference to ConfigReference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/config.go#L55)  

```go
func ConfigReferencesFromGRPC(s []*swarmapi.ConfigReference) []*swarmtypes.ConfigReference
```

---

### ConfigSpecToGRPC

ConfigSpecToGRPC converts Config to a grpc Config.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/config.go#L35)  

```go
func ConfigSpecToGRPC(s swarmtypes.ConfigSpec) swarmapi.ConfigSpec
```

---

### GenericResourcesFromGRPC

GenericResourcesFromGRPC converts a GRPC GenericResource to a GenericResource

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/service.go#L379)  

```go
func GenericResourcesFromGRPC(genericRes []*swarmapi.GenericResource) []types.GenericResource
```

---

### GenericResourcesToGRPC

GenericResourcesToGRPC converts a GenericResource to a GRPC GenericResource

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/service.go#L442)  

```go
func GenericResourcesToGRPC(genericRes []types.GenericResource) []*swarmapi.GenericResource
```

---

### IsIngressNetwork

IsIngressNetwork check if the swarm network is an ingress network

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/network.go#L235)  

```go
func IsIngressNetwork(n *swarmapi.Network) bool
```

---

### IsolationFromGRPC

IsolationFromGRPC converts a swarm api container isolation to a moby isolation representation

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/container.go#L524)  

```go
func IsolationFromGRPC(i swarmapi.ContainerSpec_Isolation) container.Isolation
```

---

### MergeSwarmSpecToGRPC

MergeSwarmSpecToGRPC merges a Spec with an initial grpc ClusterSpec

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/swarm.go#L90)  
**Added in:** v1.13.0

```go
func MergeSwarmSpecToGRPC(s types.Spec, spec swarmapi.ClusterSpec) (swarmapi.ClusterSpec, error)
```

---

### NodeFromGRPC

NodeFromGRPC converts a grpc Node to a Node.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/node.go#L13)  

```go
func NodeFromGRPC(n swarmapi.Node) types.Node
```

---

### NodeSpecToGRPC

NodeSpecToGRPC converts a NodeSpec to a grpc NodeSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/node.go#L93)  

```go
func NodeSpecToGRPC(s types.NodeSpec) (swarmapi.NodeSpec, error)
```

---

### SecretFromGRPC

SecretFromGRPC converts a grpc Secret to a Secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/secret.go#L10)  
**Added in:** v1.13.0

```go
func SecretFromGRPC(s *swarmapi.Secret) swarmtypes.Secret
```

---

### SecretReferencesFromGRPC

SecretReferencesFromGRPC converts a slice of grpc SecretReference to SecretReference

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/secret.go#L57)  
**Added in:** v1.13.0

```go
func SecretReferencesFromGRPC(s []*swarmapi.SecretReference) []*swarmtypes.SecretReference
```

---

### SecretSpecToGRPC

SecretSpecToGRPC converts Secret to a grpc Secret.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/secret.go#L36)  
**Added in:** v1.13.0

```go
func SecretSpecToGRPC(s swarmtypes.SecretSpec) swarmapi.SecretSpec
```

---

### ServiceFromGRPC

ServiceFromGRPC converts a grpc Service to a Service.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/service.go#L25)  

```go
func ServiceFromGRPC(s swarmapi.Service) (types.Service, error)
```

---

### ServiceSpecToGRPC

ServiceSpecToGRPC converts a ServiceSpec to a grpc ServiceSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/service.go#L157)  

```go
func ServiceSpecToGRPC(s types.ServiceSpec) (swarmapi.ServiceSpec, error)
```

---

### SwarmFromGRPC

SwarmFromGRPC converts a grpc Cluster to a Swarm.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/swarm.go#L14)  

```go
func SwarmFromGRPC(c swarmapi.Cluster) types.Swarm
```

---

### SwarmPluginGetter

SwarmPluginGetter adapts a plugingetter.PluginGetter to a Swarmkit plugin.Getter.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/pluginadapter.go#L9)  

```go
func SwarmPluginGetter(pg plugingetter.PluginGetter) plugin.Getter
```

---

### SwarmSpecToGRPC

SwarmSpecToGRPC converts a Spec to a grpc ClusterSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/swarm.go#L85)  

```go
func SwarmSpecToGRPC(s types.Spec) (swarmapi.ClusterSpec, error)
```

---

### TaskFromGRPC

TaskFromGRPC converts a grpc Task to a Task.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/task.go#L12)  

```go
func TaskFromGRPC(t swarmapi.Task) (types.Task, error)
```

---

### VolumeCreateToGRPC

VolumeCreateToGRPC takes a VolumeCreateBody and outputs the matching
swarmapi VolumeSpec.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/volume.go#L137)  

```go
func VolumeCreateToGRPC(volume *volumetypes.CreateOptions) *swarmapi.VolumeSpec
```

---

### VolumeFromGRPC

VolumeFromGRPC converts a swarmkit api Volume object to a docker api Volume
object

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/daemon/cluster/convert/volume.go#L11)  

```go
func VolumeFromGRPC(v *swarmapi.Volume) volumetypes.Volume
```

---

## Types

This section is empty.

