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

package infradriver

import (
	"net"

	v1 "github.com/sealerio/sealer/types/api/v1"
)

// InfraDriver treat the entire cluster as an operating system kernel,
// interface function here is the target system call.
type InfraDriver interface {
	GetHostIPList() []net.IP

	GetHostIPListByRole(role string) []net.IP

	GetHostsPlatform(hosts []net.IP) (map[v1.Platform][]net.IP, error)

	//GetHostEnv return merged env with host env and cluster env.
	GetHostEnv(host net.IP) map[string]interface{}

	//GetClusterEnv return cluster.spec.env as map[string]interface{}
	GetClusterEnv() map[string]interface{}

	//GetClusterName ${clusterName}
	GetClusterName() string

	//GetClusterImageName ${cluster image Name}
	GetClusterImageName() string

	//GetClusterLaunchCmds ${user-defined launch command}
	GetClusterLaunchCmds() []string

	//GetClusterRootfsPath /var/lib/sealer/data/${clusterName}/rootfs
	GetClusterRootfsPath() string

	// GetClusterBasePath /var/lib/sealer/data/${clusterName}
	GetClusterBasePath() string

	// Execute use eg.Go to execute shell cmd concurrently
	Execute(hosts []net.IP, f func(host net.IP) error) error

	// Copy local files to remote host
	// scp -r /tmp root@192.168.0.2:/root/tmp => Copy("192.168.0.2","tmp","/root/tmp")
	// need check md5sum
	Copy(host net.IP, localFilePath, remoteFilePath string) error
	// CopyR copy remote host files to localhost
	CopyR(host net.IP, remoteFilePath, localFilePath string) error
	// CmdAsync exec command on remote host, and asynchronous return logs
	CmdAsync(host net.IP, cmd ...string) error
	// Cmd exec command on remote host, and return combined standard output and standard error
	Cmd(host net.IP, cmd string) ([]byte, error)
	// CmdToString exec command on remote host, and return spilt standard output and standard error
	CmdToString(host net.IP, cmd, spilt string) (string, error)

	// IsFileExist check remote file exist or not
	IsFileExist(host net.IP, remoteFilePath string) (bool, error)
	// IsDirExist Remote file existence returns true, nil
	IsDirExist(host net.IP, remoteDirPath string) (bool, error)

	// GetPlatform Get remote platform
	GetPlatform(host net.IP) (v1.Platform, error)

	GetHostName(host net.IP) (string, error)
	// Ping Ping remote host
	Ping(host net.IP) error
	// SetHostName add or update host name on host
	SetHostName(host net.IP, hostName string) error
	// SetLvsRule add or update host name on host
	//SetLvsRule(host net.IP, hostName string) error
}
