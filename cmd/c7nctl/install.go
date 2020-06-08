// Copyright © 2018 VinkDong <dong@wenqi.us>
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

package main

import (
	"github.com/choerodon/c7nctl/pkg/action"
	"github.com/spf13/cobra"
	"io"
)

// TODO REMOVE
var (
	ConfigFile   string
	ResourceFile string
)

const installDesc = `
install [command] [flags]
`

// installCmd represents the resource command
func newInstallCmd(cfg *action.Configuration, out io.Writer, args []string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install choerodon、kubernetes or component.",
		Long:  installC7nDesc,
	}

	cmd.AddCommand(
		newInstallC7nCmd(cfg, out, args),
		newInstallK8sCmd(out, args),
		newInstallComponentCmd(cfg, args),
	)

	return cmd
}