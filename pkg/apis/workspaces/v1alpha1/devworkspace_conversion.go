package v1alpha1

import (
	"errors"
	"github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	conv "sigs.k8s.io/controller-runtime/pkg/conversion"
	apimachineryconv "k8s.io/apimachinery/pkg/conversion"
)

// Spokes for conversion have to satisfy the Convertible interface.
var _ conv.Convertible = (*DevWorkspace)(nil)

func (src *DevWorkspace) ConvertTo(dstRaw conv.Hub) error {
	dst := dstRaw.(*v1alpha2.DevWorkspace)

	// Basic conversions of metadata and status
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.WorkspaceId = src.Status.WorkspaceId
	dst.Status.IdeUrl = src.Status.IdeUrl
	dst.Status.Phase = v1alpha2.WorkspacePhase(src.Status.Phase)


	return nil
}

func (dst *DevWorkspace) ConvertFrom(srcRaw conv.Hub) error {
	return errors.New("Unimplemented")
}

func autoConvert_v1alpha1_Component_To_v1alpha2_Component(in *Component, out *v1alpha2.Component, s apimachineryconv.Scope) error {
	return nil
}