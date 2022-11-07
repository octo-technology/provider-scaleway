/*
Copyright 2021 Upbound Inc.
*/

package instance_server

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("scaleway_instance_server", func(r *ujconfig.Resource) {
		// And other overrides.
	})
}
