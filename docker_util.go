package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type imagePullStatus struct {
	Status         string      `json:"status"`
	ProgressDetail interface{} `json:"progressDetail"`
	ID             string      `json:"Id"`
}

func ImageExists(ctx context.Context, clnt *client.Client, imageTag string) bool {
	filterArgs := filters.NewArgs()
	filterArgs.Add("dangling", "false")
	images, err := clnt.ImageList(ctx, types.ImageListOptions{
		Filters: filterArgs,
	})
	CheckError(err)

	for _, image := range images {
		if matchRepoTag(image, imageTag) {
			return true
		}
	}

	return false
}

func matchRepoTag(sum types.ImageSummary, tagToMatch string) bool {
	for _, tag := range sum.RepoTags {
		if tag == tagToMatch {
			return true
		}
	}
	return false
}

func PullImage(clnt *client.Client, image string) {
	imagePath := "docker.io/" + image

	log.Printf("Pull image from %s\n", imagePath)
	reader, err := clnt.ImagePull(context.Background(), imagePath, types.ImagePullOptions{})
	defer reader.Close()
	CheckError(err)

	scanner := bufio.NewScanner(reader)
	status := imagePullStatus{}
	var lastStatus string

	for scanner.Scan() {
		json.Unmarshal(scanner.Bytes(), &status)

		if status.Status != lastStatus {
			if *verbose {
				log.Printf("Pull image[%s]: %s\n", status.ID, status.Status)
			}
		}
		lastStatus = status.Status
	}

	log.Println("Image pulled successfully")
}

func CreateFlywayContainer(ctx context.Context, clnt *client.Client, image string, flywayCommand string) string {
	confPath, err := filepath.Abs(fmt.Sprintf("./conf/%s.conf", *profile))
	CheckError(err)

	sqlPath, err := filepath.Abs("./sql")
	CheckError(err)

	container, err := clnt.ContainerCreate(context.Background(),
		&container.Config{
			Image:        image,
			Cmd:          []string{flywayCommand},
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
		&container.HostConfig{
			AutoRemove:  true,
			NetworkMode: "host",
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: confPath,
					Target: "/flyway/conf/flyway.conf",
				},
				{
					Type:   mount.TypeBind,
					Source: sqlPath,
					Target: "/flyway/sql",
				},
			},
		},
		&network.NetworkingConfig{}, "flyway-migration")
	CheckError(err)
	return container.ID
}

func AttachContainer(ctx context.Context, clnt *client.Client, cid string) {
	resp, err := clnt.ContainerAttach(ctx, cid,
		types.ContainerAttachOptions{
			Stream: true,
			Stdout: true,
			Stderr: true,
			Logs:   true,
		},
	)

	CheckError(err)

	go io.Copy(os.Stdout, resp.Reader)
}
