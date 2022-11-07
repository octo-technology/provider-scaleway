/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/octo-technology/provider-scaleway/config/instance_server"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "scaleway"
	modulePath     = "github.com/octo-technology/provider-scaleway"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		instance_server.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
