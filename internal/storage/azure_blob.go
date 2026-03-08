//go:build azure

package storage

// Azure Blob support is intentionally behind a build tag to avoid pulling Azure SDK
// dependencies into S3-only deployments.
//
// To enable:
//   go build -tags azure ./cmd/server
//
// Implementations can be added later without impacting default builds.
