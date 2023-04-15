package contmod

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type OgreDockConfig struct {
	Client	*client.Client
	HostCon	*container.HostConfig
	NetCon	*network.NetworkingConfig
}

var OConfig OgreDockConfig;

func initializeHostConfig(port string) *container.HostConfig{
	newport, err := nat.NewPort("tcp", port)
	if err != nil {
		panic(err)
	}

	hostConfig := &container.HostConfig{
			PortBindings: nat.PortMap{
				newport: []nat.PortBinding{
					{
						HostIP: "0.0.0.0",
						HostPort: port,
					},
				},
			},
			RestartPolicy: container.RestartPolicy{
				Name: "always",
			},
	}

	return hostConfig;
}

func initializeNetworkConfig(addr string , gateway string, net string) *network.NetworkingConfig{
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	endpointConfig := &network.EndpointSettings{
		Gateway: gateway,
		IPAddress: addr,
	}
	networkConfig.EndpointsConfig[net] = endpointConfig

	return networkConfig
}

func ListContainers(){
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.Names[0])
	}
}

func GetContainers() []types.Container{
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true,})
	if err != nil {
		panic(err)
	}

	return containers
}

func StartContainer(contID string) error{
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	err2 := cli.ContainerStart(context.Background(), contID, types.ContainerStartOptions{})
	if err2 != nil{
		panic(err2)
	}

	return err2
}

func StopContainer(contID string) error{
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	timeout := 0
	err2 := cli.ContainerStop(context.Background(), contID, container.StopOptions{Timeout: &timeout})
	//err2 := cli.ContainerPause(context.Background(), contID)
	if err2 != nil{
		panic(err2)
	}

	return err2
}

func CreateContainer(name string, net string, img string, ip string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	config := &container.Config{
		Image: img,
	}

	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			net: &network.EndpointSettings{
			},
		},
	}

	cont, err2 := cli.ContainerCreate(context.Background(), config, nil, networkConfig, nil, name)
	if err2 != nil {
		panic(err2)
	}

	return cont.ID, err2;
}

func DestroyContainer(contID string) error{
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	err2 := cli.ContainerRemove(context.Background(), contID, types.ContainerRemoveOptions{Force: true,})
	if err2 != nil {
		panic(err2)
	}

	return err2
}
