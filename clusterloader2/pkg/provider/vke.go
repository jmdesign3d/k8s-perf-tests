/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"

	clientset "k8s.io/client-go/kubernetes"
	prom "k8s.io/perf-tests/clusterloader2/pkg/prometheus/clients"
)

type VKEProvider struct {
	features Features
}

func NewVKEProvider(_ map[string]string) Provider {
	return &VKEProvider{
		features: Features{
			SupportProbe:                        true,
			SupportSSHToMaster:                  false,
			SupportImagePreload:                 true,
			SupportEnablePrometheusServer:       true,
			SupportGrabMetricsFromKubelets:      true,
			SupportAccessAPIServerPprofEndpoint: true,
			SupportMetricsServerMetrics:         true,
			SupportResourceUsageMetering:        true,
			ShouldScrapeKubeProxy:               true,
		},
	}
}

func (p *VKEProvider) Name() string {
	return VKEName
}

func (p *VKEProvider) Features() *Features {
	return &p.features
}

func (p *VKEProvider) GetComponentProtocolAndPort(componentName string) (string, int, error) {
	return getComponentProtocolAndPort(componentName)
}

func (p *VKEProvider) GetConfig() Config {
	return Config{}
}

func (p *VKEProvider) RunSSHCommand(cmd, host string) (string, string, int, error) {
	return "", "", 0, fmt.Errorf("SSH not supported for VKE")
}

func (p *VKEProvider) Metadata(_ clientset.Interface) (map[string]string, error) {
	return nil, nil
}

func (p *VKEProvider) GetManagedPrometheusClient() (prom.Client, error) {
	return nil, nil
}
