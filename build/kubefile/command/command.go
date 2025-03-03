// Copyright © 2022 Alibaba Group Holding Ltd.
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

package command

const (

	// the following commands are kubefile-spec commands
	// we recommender users to use these commands to pack their
	// demands for applications.
	App    = "app"
	Launch = "launch"
	Cmds   = "cmds"

	Label      = "label"
	Maintainer = "maintainer"

	// the following commands are the intenal implementations for kube commands
	Add  = "add"
	Arg  = "arg"
	Copy = "copy"
	From = "from"
	Run  = "run"
)

// SupportedCommands is list of all Kubefile commands
var SupportedCommands = map[string]struct{}{
	Add:        {},
	Arg:        {},
	Copy:       {},
	From:       {},
	Label:      {},
	Maintainer: {},
	Run:        {},
	App:        {},
	Launch:     {},
	Cmds:       {},
}
