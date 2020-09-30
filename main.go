package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const FLYWAY_IAMGE = "flyway/flyway:latest"

var verbose *bool
var profile *string

func main() {
	profile = flag.String("p", "local", "Configuration profile, 'local' is used by default")
	verbose = flag.Bool("v", false, "Prints outputs verbosely")
	flag.Parse()

	PrintTitle()
	log.Printf("Enabled profile: '%s'\n", *profile)

	argLen := len(os.Args)
	flywayCommand := os.Args[argLen-1]

	clnt, err := client.NewEnvClient()
	CheckError(err)

	ctx := context.Background()

	if !ImageExists(ctx, clnt, FLYWAY_IAMGE) {
		PullImage(clnt, FLYWAY_IAMGE)
	}

	containerID := CreateFlywayContainer(ctx, clnt, FLYWAY_IAMGE, flywayCommand)

	log.Printf("Created container, ID: %s\n", containerID)

	AttachContainer(ctx, clnt, containerID)

	err = clnt.ContainerStart(ctx,
		containerID,
		types.ContainerStartOptions{})

	CheckError(err)

	_, err = clnt.ContainerWait(ctx, containerID)

	CheckError(err)
}
