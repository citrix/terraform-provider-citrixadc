package vpneula

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpneulaResourceModel describes the resource data model.
type VpneulaResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *VpneulaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpneula resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the eula",
			},
		},
	}
}

func vpneulaGetThePayloadFromtheConfig(ctx context.Context, data *VpneulaResourceModel) vpn.Vpneula {
	tflog.Debug(ctx, "In vpneulaGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpneula := vpn.Vpneula{}
	if !data.Name.IsNull() {
		vpneula.Name = data.Name.ValueString()
	}

	return vpneula
}

func vpneulaSetAttrFromGet(ctx context.Context, data *VpneulaResourceModel, getResponseData map[string]interface{}) *VpneulaResourceModel {
	tflog.Debug(ctx, "In vpneulaSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
