package vpnglobal_vpnnexthopserver_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalVpnnexthopserverBindingResourceModel describes the resource data model.
type VpnglobalVpnnexthopserverBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nexthopserver          types.String `tfsdk:"nexthopserver"`
}

func (r *VpnglobalVpnnexthopserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnnexthopserver_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"nexthopserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the next hop server bound to vpn global.",
			},
		},
	}
}

func vpnglobal_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalVpnnexthopserverBindingResourceModel) vpn.Vpnglobalvpnnexthopserverbinding {
	tflog.Debug(ctx, "In vpnglobal_vpnnexthopserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_vpnnexthopserver_binding := vpn.Vpnglobalvpnnexthopserverbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_vpnnexthopserver_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Nexthopserver.IsNull() {
		vpnglobal_vpnnexthopserver_binding.Nexthopserver = data.Nexthopserver.ValueString()
	}

	return vpnglobal_vpnnexthopserver_binding
}

func vpnglobal_vpnnexthopserver_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnnexthopserverBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnnexthopserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnnexthopserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["nexthopserver"]; ok && val != nil {
		data.Nexthopserver = types.StringValue(val.(string))
	} else {
		data.Nexthopserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Nexthopserver.ValueString())

	return data
}
