package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		DisableCompression: true,
	}
	httpcli := &http.Client{Transport: tr}
	cli, err := client.NewClient("tcp://localhost:2375", "1.21", httpcli, nil)

	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(context.Background())

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	

	for _, image := range images {
		date := time.Unix(image.Created,0)
		fmt.Printf("%s %s %s\n", image.ID[:10], image.RepoTags[0], date.Format(time.UnixDate) )
	}
}
