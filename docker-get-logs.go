package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	
	out, err := cli.ContainerLogs(ctx, "3433cafbdb4cf0d970584ac977e78d85179ad3beb761544fdfb9b04cf78f2afb", options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}