/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package mg

import (
	"github.com/spf13/cobra"
	impl "github.com/wso2/product-apim-tooling/import-export-cli/impl/mg"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

const (
	addEnvCmdShortDesc = "Add Environment to Config file"
	addEnvCmdLongDesc  = `Add new environment and its related endpoints to the config file`
)
const addEnvCmdExamples = utils.ProjectName + " " + mgCmdLiteral + " " + addCmdLiteral + " " + envCmdLiteral +
	" prod --adapter https://localhost:9843 " +

	"\n\nNOTE: The flag --adapter (-a) is mandatory and it has to specify the microgateway adapter" +
	" url."

// addEnvCmd represents the addEnv command
var AddEnvCmd = &cobra.Command{
	Use:     envCmdLiteral,
	Short:   addEnvCmdShortDesc,
	Long:    addEnvCmdLongDesc,
	Example: addEnvCmdExamples,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Logln(utils.LogPrefixInfo + envCmdLiteral + " called")

		envToBeAdded := args[0]

		envEndpoints := new(utils.MgwEndpoints)
		envEndpoints.AdapterEndpoint = mgwAdapterHost + impl.DefaultMgwAdapterEndpointSuffix
		err := impl.AddEnv(envToBeAdded, envEndpoints)
		if err != nil {
			utils.HandleErrorAndExit("Error adding environment", err)
		}
	},
}

// init using Cobra
func init() {
	AddCmd.AddCommand(AddEnvCmd)

	AddEnvCmd.Flags().StringVarP(&mgwAdapterHost, "adapter", "a", "", "The adapter host url with port")

	_ = AddEnvCmd.MarkFlagRequired("adapter")
}
