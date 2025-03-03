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

package imagedistributor

import (
	"net"

	v1 "github.com/sealerio/sealer/types/api/v1"
)

type Distributor interface {
	// DistributeRootfs each files under mounted cluster image directory to target hosts.
	DistributeRootfs(hosts []net.IP, rootfsPath string) error
	// DistributeRegistry each files under registry directory to target hosts.
	DistributeRegistry(deployHost net.IP, dataDir string) error
	// Restore will do some clean works via infra driver, like delete rootfs.
	Restore(targetDir string, hosts []net.IP) error
}

type Mounter interface {
	Mount(imageName string, platform v1.Platform) (string, error)
	Umount(dir string) error
}
