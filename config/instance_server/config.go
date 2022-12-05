/*
Copyright 2021 Upbound Inc.
*/

package instance_server

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	//p.AddResourceConfigurator("scaleway_instance_server", func(r *ujconfig.Resource) {
	//	// And other overrides.
	//})
	p.AddResourceConfigurator("scaleway_instance_server", func (r *ujconfig.Resource) {
		r.References["private_network.pn_id"] = ujconfig.Reference { //reference vers le param de la ressource
			Type: "github.com/octo-technology/provider-scaleway/apis/vpc/v1alpha1.PrivateNetwork",
		}
	})
}
