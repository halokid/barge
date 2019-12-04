package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
)


type cPort struct {
	IP 							string
	PrivatePort		 	int
	PublicPort			int
	Type						string
}

// Container is a Docker container
type Container struct {
	ID    		string
	Image 		string
	State 		string
	Ports			[]cPort
	Names 		[]string
}


func (p Portainer) getContainersForEndpoint(endpoint Endpoint) []Container {
	// 获取endpoint的容器信息
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/containers/json")
	logx.DebugPrint("getContainersForEndpoint output --------------", output)

	containers := make([]Container, 0)

	json.Unmarshal([]byte(output), &containers)
	logx.DebugPrint("getContainersForEndpoint containers --------------", containers)

	return containers
}

func (p Portainer) populateContainersForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Containers = p.getContainersForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func printContainersForEndpoint(endpoint Endpoint) {
	fmt.Println(endpoint.Name, endpoint.ID, "容器列表")
	fmt.Println("----")

	for _, c := range endpoint.Containers {
		fmt.Println("ID: " + c.ID[0:12] + ", Name:", c.Names, "Port:", c.Ports, ", image: " + c.Image)
	}
	fmt.Println("----")
}
