package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	. "github.com/halokid/ColorfulRabbit"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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


func (p *Portainer) PullImage(eId int, imageName string) bool {
	// pull image
	data := make(map[string]interface{})
	data["fromImage"] = imageName
	data["tag"] = "latest"
	pathImg := strings.Replace(imageName, "/", "%2F", -1)
	logx.DebugPrint("pathImg -----------", pathImg)
	path := "endpoints/" + cast.ToString(eId) + "/docker/images/create?fromImage=" + pathImg +
									"&tag=latest"
	logx.DebugPrint("PullImage path --------------", path)
	rspErr := p.Post(data, path)
	return rspErr
}

func (p *Portainer) CreateCt(eId int, imgName string, cName string, cPort string,
															hostPort string) (string, error) {
	// create container， 返回容器id， 创建了之后，还需要start才能运行
	//data := make(map[string]interface{})
	data := makeCtMap(imgName, cName, cPort, hostPort)
	logx.DebugPrint("CreateCt data ---------- ", data)
	path := p.URL + `/endpoints/` + cast.ToString(eId) + `/docker/containers/create?name=` + cName
	logx.DebugPrint("CreateCt path ----------------- ", path)
	js := []byte(data)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(js))
	bearerHeader := "Bearer " + p.token
	req.Header.Set("Authorization", bearerHeader)
	req.Header.Set("Content-Type", "application/json")
	CheckError(err, "----- CreateCt set req json err")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err, "------ CreateCt request error")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	logx.DebugPrint("CreateCt resp ----------- ", string(body))

	bdJs, err := simplejson.NewJson(body)
	CheckError(err)
	return bdJs.Get("Id").MustString(), err
}


func (p *Portainer) StartCt(cId string) bool {
	// 启动容器
	data := make(map[string]interface{})
	path := "endpoints/11/docker/containers/" + cId + "/start"
	_ = p.Post(data, path)
	return true
}




