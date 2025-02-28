package main

import (
	"context"
	"io"
	"os"

_	"github.com/docker/docker/api/types/container"
        "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
_	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", image.PullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)
}
