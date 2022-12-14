/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/octo-technology/provider-scaleway/apis/vpc/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateNetwork); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateNetwork[i3].PnID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PrivateNetwork[i3].PnIDRef,
			Selector:     mg.Spec.ForProvider.PrivateNetwork[i3].PnIDSelector,
			To: reference.To{
				List:    &v1alpha1.PrivateNetworkList{},
				Managed: &v1alpha1.PrivateNetwork{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateNetwork[i3].PnID")
		}
		mg.Spec.ForProvider.PrivateNetwork[i3].PnID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateNetwork[i3].PnIDRef = rsp.ResolvedReference

	}

	return nil
}
