package metadata

import (
	"github.com/docker/infrakit/pkg/spi/metadata"
	"github.com/docker/infrakit/pkg/types"
)

// ListRequest is the rpc wrapper for request parameters to List
type ListRequest struct {
	Path types.Path
}

// ListResponse is the rpc wrapper for the results of List
type ListResponse struct {
	Nodes []string
}

// GetRequest is the rpc wrapper of the params to Get
type GetRequest struct {
	Path types.Path
}

// GetResponse is the rpc wrapper of the result of Get
type GetResponse struct {
	Value *types.Any
}

// ChangesRequest is the rpc wrapper of the params to Changes
type ChangesRequest struct {
	Changes []metadata.Change
}

// ChangesResponse is the rpc wrapper of the params to Changes
type ChangesResponse struct {
	Original *types.Any
	Proposed *types.Any
	Cas      string
}

// CommitRequest is the rpc wrapper of the params to Commit
type CommitRequest struct {
	Proposed *types.Any
	Cas      string
}

// CommitResponse is the rpc wrapper of the params to Commit
type CommitResponse struct {
}
