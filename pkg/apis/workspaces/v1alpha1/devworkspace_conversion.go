package v1alpha1

import (
	"errors"
	"github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	"k8s.io/apimachinery/pkg/conversion"
	runtimeconversion "sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Spokes for conversion have to satisfy the Convertible interface.
var _ runtimeconversion.Convertible = (*DevWorkspace)(nil)

func (src *DevWorkspace) ConvertTo(dstRaw runtimeconversion.Hub) error {
	dst := dstRaw.(*v1alpha2.DevWorkspace)

	// Basic conversions of metadata and status
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.WorkspaceId = src.Status.WorkspaceId
	dst.Status.IdeUrl = src.Status.IdeUrl
	dst.Status.Phase = v1alpha2.WorkspacePhase(src.Status.Phase)


	return nil
}

func (dst *DevWorkspace) ConvertFrom(srcRaw runtimeconversion.Hub) error {
	return errors.New("unimplemented")
}

func Convert_v1alpha1_Component_To_v1alpha2_Component(in *Component, out *v1alpha2.Component, s conversion.Scope) error {
	
}

func Convert_v1alpha2_Component_To_v1alpha1_Component(in *v1alpha2.Component, out *Component, s conversion.Scope) error {

}

