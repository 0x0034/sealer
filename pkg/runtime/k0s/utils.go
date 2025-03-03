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

package k0s

import (
	"fmt"
	"net"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/sealerio/sealer/common"
	"github.com/sealerio/sealer/pkg/clustercert/cert"
	"github.com/sealerio/sealer/utils/exec"
	osi "github.com/sealerio/sealer/utils/os"
	"github.com/sealerio/sealer/utils/ssh"

	"github.com/pkg/errors"
)

const WaitingFork0sServiceStartTimes = 5

func GenerateRegistryCert(registryCertPath string, baseName string) error {
	regCertConfig := cert.CertificateDescriptor{
		CommonName:   baseName,
		DNSNames:     []string{baseName},
		Organization: []string{common.ExecBinaryFileName},
		Year:         100,
	}
	if baseName != SeaHub {
		regCertConfig.DNSNames = append(regCertConfig.DNSNames, SeaHub)
	}

	caGenerator := cert.NewAuthorityCertificateGenerator(regCertConfig)
	caCert, caKey, err := caGenerator.Generate()
	if err != nil {
		return fmt.Errorf("unable to generate %s cert: %v", baseName, err)
	}

	// write cert file to disk
	err = cert.NewCertificateFileManger(registryCertPath, baseName).Write(caCert, caKey)
	if err != nil {
		return fmt.Errorf("unable to save %s cert: %v", baseName, err)
	}

	return nil
}

func FetchKubeconfigAndGetKubectl(ssh ssh.Interface, host net.IP, rootfs string) error {
	// fetch the cluster kubeconfig
	err := ssh.CopyR(host, path.Join(common.DefaultKubeConfigDir(), "config"), DefaultAdminConf)
	if err != nil {
		return errors.Wrap(err, "failed to copy kubeconfig")
	}

	if !osi.IsFileExist(common.DefaultKubectlPath) {
		err = osi.RecursionCopy(filepath.Join(rootfs, "bin/kubectl"), common.DefaultKubectlPath)
		if err != nil {
			return err
		}
		err = exec.Cmd("chmod", "+x", common.DefaultKubectlPath)
		if err != nil {
			return errors.Wrap(err, "failed to chmod a+x kubectl")
		}
	}
	return nil
}

func (k *Runtime) WaitK0sReady(ssh ssh.Interface, host net.IP) error {
	times := WaitingFork0sServiceStartTimes
	for {
		times--
		if times == 0 {
			break
		}
		time.Sleep(time.Second * 2)
		bytes, err := ssh.Cmd(host, "k0s status")
		if err != nil {
			return err
		}
		// k0s status return: `Process ID: xxx` when it started successfully, or return: `connect failed`,
		// so we use field `Process` whether contains in string(bytes) to verify if k0s service started successfully.
		if strings.Contains(string(bytes), "Process") {
			return nil
		}
	}
	return errors.New("failed to start k0s: failed to get k0s status after 10 seconds")
}
