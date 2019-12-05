package cmd

import (
  "encoding/json"
  "fmt"
  "testing"
)

func TestMarshalCts(t *testing.T) {
  //s := `{"Id": "xxx", "Image": "yyy", "State": "zzzz", "Ports": [{"IP": "8.8.8.8"]`
  output := `[{"Command":"/jvmIdx","Created":1575359495,"HostConfig":{"NetworkMode":"default"},"Id":"a3bd5483b93ded1401f74be2523b1ed3e949cf6c719a3fe8341a84fda559179f","Image":"10.1.1.9:5000/jvmidx:latest","ImageID":"sha256:dfa471cd036d3ea2d17d229050993f92b66cb43f268ab62a9edabf040a8ca57a","Labels":{},"Mounts":[{"Destination":"/etc/localtime","Mode":"ro","Propagation":"rprivate","RW":false,"Source":"/etc/localtime","Type":"bind"}],"Names":["/jvmIdx-test"],"NetworkSettings":{"Networks":{"bridge":{"Aliases":null,"DriverOpts":null,"EndpointID":"a7879b0a54b3047d2cc837523809e788533b83b2ee2e63682e58dc30b0fca590","Gateway":"172.17.0.1","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"IPAMConfig":null,"IPAddress":"172.17.0.3","IPPrefixLen":16,"IPv6Gateway":"","Links":null,"MacAddress":"02:42:ac:11:00:03","NetworkID":"92d9381add3ad16bd3652a8cae697b765885c963796cb1d4047ccc68a5078348"}}},"Ports":[{"IP":"0.0.0.0","PrivatePort":8080,"PublicPort":8089,"Type":"tcp"}],"State":"running","Status":"Up 19 hours"}]`
  containers := make([]Container, 0)

  json.Unmarshal([]byte(output), &containers)

  fmt.Println(containers)
}

func TestPortainer_PullImage(t *testing.T) {
  //return
  //p := NewPortainer()
  //data := make(map[string]interface{})
  //data["fromImage"] = "10.1.1.9:5000/micro_demo"
  //data["tag"] = "latest"
  //path := "http://10.1.1.40:9000/api/endpoints/11/docker/images/create?fromImage=10.1.1.9:5000%2Fmicro_demo&tag=latest"
  //p.PullImage(data, path)

  p := NewPortainer()
  pullRes := p.PullImage(11, "10.1.1.9:5000/micro_demo")
  fmt.Println(pullRes)
}

func TestPortainer_CreateCt(t *testing.T) {
  p := NewPortainer()
  cId, err := p.CreateCt(11, "xx-dev")
  fmt.Println(cId, err)
}




