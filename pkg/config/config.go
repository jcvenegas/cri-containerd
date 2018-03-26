/*
Copyright 2017 The Kubernetes Authors.

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

package config

import "github.com/containerd/containerd"

// Runtime  struct to contain the type(ID), engine, and root variables for a default and a privileged runtime
type Runtime struct {
	//Type is the runtime type to use in containerd e.g. io.containerd.runtime.v1.linux
	Type string `toml:"runtime_type" json:"runtimeType,omitempty"`
	// Engine is the name of the runtime engine used by containerd.
	Engine string `toml:"runtime_engine" json:"runtimeEngine,omitempty"`
	// Root is the directory used by containerd for runtime state.
	Root string `toml:"runtime_root" json:"runtimeRoot,omitempty"`
}

// ContainerdConfig contains toml config related to containerd
type ContainerdConfig struct {
	// Snapshotter is the snapshotter used by containerd.
	Snapshotter string `toml:"snapshotter" json:"snapshotter,omitempty"`
	// DefaultRuntime is the runtime to use in containerd.
	DefaultRuntime Runtime `toml:"default_runtime" json:"defaultRuntime,omitempty"`
	// PrivilegedRuntime is a non-secure runtime used only to run trusted workloads on it
	PrivilegedRuntime Runtime `toml:"privileged_runtime" json:"privilegedRuntime,omitempty"`
}

// CniConfig contains toml config related to cni
type CniConfig struct {
	// NetworkPluginBinDir is the directory in which the binaries for the plugin is kept.
	NetworkPluginBinDir string `toml:"bin_dir" json:"binDir,omitempty"`
	// NetworkPluginConfDir is the directory in which the admin places a CNI conf.
	NetworkPluginConfDir string `toml:"conf_dir" json:"confDir,omitempty"`
}

// Mirror contains the config related to the registry mirror
type Mirror struct {
	// Endpoints are endpoints for a namespace. CRI plugin will try the endpoints
	// one by one until a working one is found.
	Endpoints []string `toml:"endpoint" json:"endpoint,omitempty"`
	// TODO (Abhi) We might need to add auth per namespace. Looks like
	// image auth information is passed by kube itself.
}

// Registry is registry settings configured
type Registry struct {
	// Mirrors are namespace to mirror mapping for all namespaces.
	Mirrors map[string]Mirror `toml:"mirrors" json:"mirrors,omitempty"`
}

// PluginConfig contains toml config related to CRI plugin,
// it is a subset of Config.
type PluginConfig struct {
	// ContainerdConfig contains config related to containerd
	ContainerdConfig `toml:"containerd" json:"containerd,omitempty"`
	// CniConfig contains config related to cni
	CniConfig `toml:"cni" json:"cni,omitempty"`
	// Registry contains config related to the registry
	Registry `toml:"registry" json:"registry,omitempty"`
	// StreamServerAddress is the ip address streaming server is listening on.
	StreamServerAddress string `toml:"stream_server_address" json:"streamServerAddress,omitempty"`
	// StreamServerPort is the port streaming server is listening on.
	StreamServerPort string `toml:"stream_server_port" json:"streamServerPort,omitempty"`
	// EnableSelinux indicates to enable the selinux support.
	EnableSelinux bool `toml:"enable_selinux" json:"enableSelinux,omitempty"`
	// SandboxImage is the image used by sandbox container.
	SandboxImage string `toml:"sandbox_image" json:"sandboxImage,omitempty"`
	// StatsCollectPeriod is the period (in seconds) of snapshots stats collection.
	StatsCollectPeriod int `toml:"stats_collect_period" json:"statsCollectPeriod,omitempty"`
	// SystemdCgroup enables systemd cgroup support.
	SystemdCgroup bool `toml:"systemd_cgroup" json:"systemdCgroup,omitempty"`
}

// Config contains all configurations for cri server.
type Config struct {
	// PluginConfig is the config for CRI plugin.
	PluginConfig
	// ContainerdRootDir is the root directory path for containerd.
	ContainerdRootDir string `json:"containerdRootDir,omitempty"`
	// ContainerdEndpoint is the containerd endpoint path.
	ContainerdEndpoint string `json:"containerdEndpoint,omitempty"`
	// RootDir is the root directory path for managing cri plugin files
	// (metadata checkpoint etc.)
	RootDir string `json:"rootDir,omitempty"`
}

// DefaultConfig returns default configurations of cri plugin.
func DefaultConfig() PluginConfig {
	return PluginConfig{
		CniConfig: CniConfig{
			NetworkPluginBinDir:  "/opt/cni/bin",
			NetworkPluginConfDir: "/etc/cni/net.d",
		},
		ContainerdConfig: ContainerdConfig{
			Snapshotter: containerd.DefaultSnapshotter,
			DefaultRuntime: Runtime{
				Type:   "io.containerd.runtime.v1.linux",
				Engine: "",
				Root:   "",
			},
			PrivilegedRuntime: Runtime{
				Type:   "io.containerd.runtime.v1.linux",
				Engine: "",
				Root:   "",
			},
		},
		StreamServerAddress: "",
		StreamServerPort:    "10010",
		EnableSelinux:       false,
		SandboxImage:        "gcr.io/google_containers/pause:3.0",
		StatsCollectPeriod:  10,
		SystemdCgroup:       false,
		Registry: Registry{
			Mirrors: map[string]Mirror{
				"docker.io": {
					Endpoints: []string{"https://registry-1.docker.io"},
				},
			},
		},
	}
}
