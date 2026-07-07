package cloudservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// cloudserviceResourceType is the NITRO resource-type string. There is no
// service.Cloudservice enum in the vendored adc-nitro-go, so the raw type string
// is used with a map payload. NO vendor/ edits.
const cloudserviceResourceType = "cloudservice"

// CloudserviceResourceModel describes the resource data model.
//
// cloudservice is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "cloudservice" object exposes no read/write properties and only the check
// action (POST /nitro/v1/config/cloudservice?action=check). The model therefore
// carries only the synthetic id.
type CloudserviceResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *CloudserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudservice resource.",
			},
		},
	}
}

// cloudserviceGetThePayloadFromthePlan builds the (empty) NITRO payload for the
// check action. cloudservice has no read/write attributes, so the payload is an
// empty map.
func cloudserviceGetThePayloadFromthePlan(ctx context.Context, data *CloudserviceResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In cloudserviceGetThePayloadFromthePlan Function")
	cloudservice := make(map[string]interface{})
	return cloudservice
}
