package nsacls6

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nsacls6ResourceModel describes the resource data model.
type Nsacls6ResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (r *Nsacls6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacls6 resource.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the acl ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
		},
	}
}

// nsacls6GetThePayloadFromthePlan builds the action payload, including only the set args.
func nsacls6GetThePayloadFromthePlan(ctx context.Context, data *Nsacls6ResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nsacls6GetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		payload["type"] = data.Type.ValueString()
	} else {
		// type defaults to CLASSIC per NITRO doc
		payload["type"] = "CLASSIC"
	}

	return payload
}
