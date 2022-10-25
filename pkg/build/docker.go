// Copyright 2022 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
	"fmt"

	"github.com/docker/docker/client"

	pb "github.com/tsuru/deploy-agent/v2/api/v1alpha1"
)

var _ pb.BuildServer = (*Docker)(nil)

type DockerOptions struct {
	TempDir string
}

func NewDocker(dc *client.Client, opts DockerOptions) *Docker {
	return &Docker{Client: dc, opts: opts}
}

type Docker struct {
	*pb.UnimplementedBuildServer
	*client.Client

	opts DockerOptions
}

func (d *Docker) Build(req *pb.BuildRequest, stream pb.Build_BuildServer) error {
	fmt.Println("Build RPC called")
	defer fmt.Println("Finishing Build RPC call")

	return stream.Send(&pb.BuildResponse{
		Data: &pb.BuildResponse_Output{Output: "Hello world!"},
	})
}
