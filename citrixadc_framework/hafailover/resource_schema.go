package hafailover

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HafailoverResourceModel describes the resource data model.
type HafailoverResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Force types.Bool   `tfsdk:"force"`
}

func (r *HafailoverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hafailover resource.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Force a failover without prompting for confirmation.",
			},
		},
	}
}

func hafailoverGetThePayloadFromtheConfig(ctx context.Context, data *HafailoverResourceModel) ha.Hafailover {
	tflog.Debug(ctx, "In hafailoverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	hafailover := ha.Hafailover{}
	if !data.Force.IsNull() {
		hafailover.Force = data.Force.ValueBool()
	}

	return hafailover
}

func hafailoverSetAttrFromGet(ctx context.Context, data *HafailoverResourceModel, getResponseData map[string]interface{}) *HafailoverResourceModel {
	tflog.Debug(ctx, "In hafailoverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["force"]; ok && val != nil {
		data.Force = types.BoolValue(val.(bool))
	} else {
		data.Force = types.BoolNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("hafailover-config")

	return data
}
