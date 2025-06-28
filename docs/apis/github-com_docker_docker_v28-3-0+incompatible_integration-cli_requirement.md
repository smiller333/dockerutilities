# Package Documentation

## Metadata

**Source:** https://pkg.go.dev/github.com/docker/docker@v28.3.0+incompatible/integration-cli/requirement

**Version:** v28.3.0+incompatible

**Generated:** 2025-06-28 21:31:01 UTC

## Constants

This section is empty.

## Variables

This section is empty.

## Functions

### Is

Is checks if the environment satisfies the requirements
for the test to run or skips the tests.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/requirement/requirement.go#L16)  

```go
func Is(t *testing.T, requirements ...Test)
```

---

## Types

### Test

Test represent a function that can be used as a requirement validation.

**Source:** [View Source](https://github.com/docker/docker/blob/v28.3.0/integration-cli/requirement/requirement.go#L12)  

```go
type Test func() bool
```

---

