package systemsignedexereport

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemsignedexereportResourceType is the exact NITRO resource name. This
// resource is not registered in the vendored service.NitroResourceType enum, so
// the literal name is used with ActOnResource.
const systemsignedexereportResourceType = "systemsignedexereport"

// SystemsignedexereportResourceModel describes the resource data model.
//
// systemsignedexereport is a ZERO-ATTRIBUTE, ACTION-ONLY (enable/disable)
// resource: the NITRO object exposes only the enable/disable actions (its single
// "message" property is read-only and only returned by the action response). The
// model therefore carries only the synthetic id.
type SystemsignedexereportResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemsignedexereportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemsignedexereport resource.",
			},
		},
	}
}

// systemsignedexereportGetThePayloadFromthePlan builds the (empty) NITRO payload
// for the enable/disable actions. This resource has no read/write attributes, so
// the payload is an empty map.
func systemsignedexereportGetThePayloadFromthePlan(ctx context.Context, data *SystemsignedexereportResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemsignedexereportGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
