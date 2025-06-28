# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/libnetwork/drivers/bridge/internal/iptabler

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:33:16 UTC

## Constants

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/iptabler/iptabler.go#L15)

```go
const (

	// DockerForwardChain contains Docker's filter-FORWARD rules.
	//
	// FIXME(robmry) - only exported because it's used to set up the jump to swarm's DOCKER-INGRESS chain.
	DockerForwardChain = "DOCKER-FORWARD"
)
```

## Variables

This section is empty.

## Functions

### NewIptabler

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/libnetwork/drivers/bridge/internal/iptabler/iptabler.go#L42)  

```go
func NewIptabler(ctx context.Context, config firewaller.Config) (firewaller.Firewaller, error)
```

---

## Types

This section is empty.

