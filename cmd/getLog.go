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
	"github.com/spf13/cobra"
	"os"
)

// listContainersCmd represents the listContainers command
var getLogCmd = &cobra.Command{
	Use:   "get",
	Short: "获取容器日志",
	Long:  `获取容器日志---`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logx.DebugPrint("需要endpoint id 和 容器id, ./tinyRabbit log get <eid> <cid>")
			os.Exit(404)
		}
		fmt.Println(args)
		eId, cId  := args[0], args[1]
		printLog(eId, cId)
	},
}

func init() {
	logCmd.AddCommand(getLogCmd)
}



