package systemautorestorefeature

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemautorestorefeatureResourceType is the exact NITRO resource name. This
// resource is not registered in the vendored service.NitroResourceType enum, so
// the literal name is used with ActOnResource.
const systemautorestorefeatureResourceType = "systemautorestorefeature"

// SystemautorestorefeatureResourceModel describes the resource data model.
//
// systemautorestorefeature is a ZERO-ATTRIBUTE, ACTION-ONLY (enable/disable)
// resource: the NITRO object exposes no read/write properties and only the
// enable/disable actions. The model therefore carries only the synthetic id.
type SystemautorestorefeatureResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemautorestorefeatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemautorestorefeature resource.",
			},
		},
	}
}

// systemautorestorefeatureGetThePayloadFromthePlan builds the (empty) NITRO
// payload for the enable/disable actions. This resource has no read/write
// attributes, so the payload is an empty map.
func systemautorestorefeatureGetThePayloadFromthePlan(ctx context.Context, data *SystemautorestorefeatureResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemautorestorefeatureGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
