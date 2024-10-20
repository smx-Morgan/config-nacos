// Copyright 2024 CloudWeGo Authors
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

package client

import (
	"github.com/kitex-contrib/config-nacos/v2/nacos"
	"github.com/kitex-contrib/config-nacos/v2/utils"

	configclient "github.com/cloudwego-contrib/cwgo-pkg/config/nacos/v2/client"
)

// NacosClientSuite nacos client config suite, configure retry timeout limit and circuitbreak dynamically from nacos.
type NacosClientSuite = configclient.NacosClientSuite

// NewSuite service is the destination service name and client is the local identity.
func NewSuite(service, client string, cli nacos.Client, opts ...utils.Option) *NacosClientSuite {
	return configclient.NewSuite(service, client, cli, opts...)
}
