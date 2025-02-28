package provider

import (
	"log"
	"io"
	"os"
	"fmt"
	"context"
	"math/rand/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
        "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
)

func resourceDockerContainer() *schema.Resource {

	return &schema.Resource {
		Description: "This resource will support create, read, update and delete containers via Terraform.",

		CreateContext: resourceDockerContainerCreate,
		ReadContext:   resourceDockerContainerRead,
		UpdateContext: resourceDockerContainerUpdate,
		DeleteContext: resourceDockerContainerDelete,

		Schema: map[string]*schema.Schema {
			"container_name": {
				Description: "Name of the docker container.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew: true,
			},
			"image_name": {
				Description: "Name of the docker image.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew: true,
			},
			"host_name": {
				Description: "Hostname of the docker container.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func pullDockerImage( imageName string ) {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)
}

func resourceDockerContainerCreate( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

    if err != nil {
        panic(err)
    }
    defer cli.Close()

    containerName := d.Get("container_name").(string)
    imageName  := d.Get("image_name").(string)
    hostName  := d.Get("host_name").(string)

   // pullDockerImage ( imageName )

    containerName = fmt.Sprintf("%v-%d", containerName, rand.IntN(99999))

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: imageName,
	Hostname: containerName,
    }, nil, nil, nil, containerName)
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
        panic(err)
    }

    d.SetId(resp.ID)
    d.Set("container_name", containerName)
    d.Set("host_name", hostName)
    d.Set("image_name", imageName )

    return nil
}

func resourceDockerContainerRead( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
    return nil
}

func resourceDockerContainerUpdate( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}

/*
func resourceDockerContainerDelete( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

    if err != nil {
        panic(err)
    }
    defer cli.Close()
	return nil

    containerID := d.Get("id").(string) 

    stopOptions := container.StopOptions{}

    if err := cli.ContainerStop(ctx, containerID, stopOptions); err != nil {
       log.Printf("Unable to stop container %s: %s", containerID, err)
    }

    removeOptions := container.RemoveOptions{RemoveVolumes: true, Force: true }

    if err := cli.ContainerRemove(ctx, containerID, removeOptions); err != nil {
       log.Printf("Unable to remove container: %s", err)
       return nil
    }
    
    return nil
}
*/
func resourceDockerContainerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics { 
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containerName := d.Get("container_name").(string)

	removeOptions := container.RemoveOptions{RemoveVolumes: true, Force: true}

	if err := cli.ContainerRemove(ctx, containerName, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return nil 
	}

	return nil
}
