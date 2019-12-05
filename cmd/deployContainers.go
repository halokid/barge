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
	"fmt"
	. "github.com/halokid/ColorfulRabbit"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

// listContainersCmd represents the listContainers command
var deployContainersCmd = &cobra.Command{
	Use:   "create",
	Short: "创建一个容器",
	Long:  `创建容器的相关操作.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logx.DebugPrint("./run <endpointId> <image_name> <container_name>" +
				" <potr:hostPort> ")
			os.Exit(404)
		}
		eId := cast.ToInt(args[0])
		imgName := args[1]
		cName := args[2]
		ports := args[3]
		cPort := strings.Split(ports, ":")[0]
		hostPort := strings.Split(ports, ":")[1]

		p := NewPortainer()
		// pull image
		pullRes := p.PullImage(eId, imgName)
		logx.DebugPrint("pullRes ----------", pullRes)
		if !pullRes {
			fmt.Println("pull容器失败")
			os.Exit(500)
		}

		time.Sleep(1 * time.Second)
		// create container
		cId, err := p.CreateCt(eId, imgName, cName, cPort, hostPort)
		logx.DebugPrint("cId -------------- ", cId)
		CheckFatal(err, " ------------ 创建容器失败")

		// start container

	},
}

func init() {
	containerCmd.AddCommand(deployContainersCmd)
}



