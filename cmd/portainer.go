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
	"github.com/spf13/viper"
)

// Portainer is an instance of Portainer
type Portainer struct {
	URL       string
	username  string
	password  string
	token     string
	verbose   bool
	Endpoints []Endpoint
}

// NewPortainer returns a new Portainer instance
func NewPortainer() Portainer {
	// 从配置文件读取这些参数
	//url := viper.GetString("portainer_url") + "/api"
	// fixme: just for test
	url :=  "00/api"

	username := viper.GetString("portainer_username")
	password := viper.GetString("portainer_password")

	portainer := Portainer{
		URL: url,
		username: username,
		password: password,
	}
	portainer.token = portainer.login()

	return portainer
}
