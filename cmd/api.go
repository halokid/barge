// Copyright © 2019 Ben Overmyer <ben@overmyer.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/halokid/ColorfulRabbit"
	"io/ioutil"
	"net/http"
)

func (p Portainer) fetch(item string) string {
	bearerHeader := "Bearer " + p.token
	requestURL := p.URL + "/" + item
	req, err := http.NewRequest("GET", requestURL, nil)

	req.Header.Set("Authorization", bearerHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (p Portainer) login() string {
	var data map[string]interface{}

	token := ""
	url := p.URL + "/auth"
	//authString := `{"Username": "` + p.username + `", "Password": "` + p.password + `"}`

	// fixme: just for test
	//url := "http://10.1.1.40:9000/api/auth"
	authString := `{"Username": "aaaaa", "Password": "aaaaaa"}`

	logx.DebugPrint(authString)

	jsonBlock := []byte(authString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBlock))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		CheckFatal(err, "-------[ERROR] auth")
		//panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &data)

	token = data["jwt"].(string)
	logx.DebugPrint("http code ---------", resp.StatusCode)
	logx.DebugPrint("token -----------------", token)

	return token
}

func (p Portainer) makePublic(resourceType string, id string) bool {
	data := map[string]interface{}{
		"Type":       resourceType,
		"Public":     true,
		"ResourceID": id,
	}
	return p.Post(data, "resource_controls")
}

func (p Portainer) Post(data map[string]interface{}, path string) bool {
	bearerHeader := "Bearer " + p.token
	requestURL := p.URL + "/" + path

	logx.DebugPrint("requestURl --------", requestURL)

	bytesData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bytesData))

	req.Header.Set("Authorization", bearerHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if p.verbose {
		fmt.Println("Sent request with data: " + string(bytesData))
		fmt.Println("Status " + resp.Status + " received from API, response was: " + string(body))
	}

	if resp.StatusCode == 200 {
		return true
	}

	logx.DebugPrint("Post rsp -----------", string(body))
	return false
}
