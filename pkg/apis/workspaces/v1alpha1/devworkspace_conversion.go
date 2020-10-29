package v1alpha1

import (
	"errors"
	"github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Spokes for conversion have to satisfy the Convertible interface.
var _ conversion.Convertible = (*DevWorkspace)(nil)

func (src *DevWorkspace) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.DevWorkspace)

	// Basic conversions of metadata and status
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.WorkspaceId = src.Status.WorkspaceId
	dst.Status.IdeUrl = src.Status.IdeUrl
	dst.Status.Phase = v1alpha2.WorkspacePhase(src.Status.Phase)


	return nil
}

func (dst *DevWorkspace) ConvertFrom(srcRaw conversion.Hub) error {
	return errors.New("Unimplemented")
}
