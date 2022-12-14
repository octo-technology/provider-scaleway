/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	server "github.com/octo-technology/provider-scaleway/internal/controller/instance/server"
	providerconfig "github.com/octo-technology/provider-scaleway/internal/controller/providerconfig"
	privatenetwork "github.com/octo-technology/provider-scaleway/internal/controller/vpc/privatenetwork"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		server.Setup,
		providerconfig.Setup,
		privatenetwork.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
