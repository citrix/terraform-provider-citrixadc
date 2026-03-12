package vpnglobal_vpnportaltheme_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalVpnportalthemeBindingResourceModel describes the resource data model.
type VpnglobalVpnportalthemeBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Portaltheme            types.String `tfsdk:"portaltheme"`
}

func (r *VpnglobalVpnportalthemeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnportaltheme_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"portaltheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the portal theme bound to vpnglobal",
			},
		},
	}
}

func vpnglobal_vpnportaltheme_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalVpnportalthemeBindingResourceModel) vpn.Vpnglobalvpnportalthemebinding {
	tflog.Debug(ctx, "In vpnglobal_vpnportaltheme_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_vpnportaltheme_binding := vpn.Vpnglobalvpnportalthemebinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_vpnportaltheme_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Portaltheme.IsNull() {
		vpnglobal_vpnportaltheme_binding.Portaltheme = data.Portaltheme.ValueString()
	}

	return vpnglobal_vpnportaltheme_binding
}

func vpnglobal_vpnportaltheme_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnportalthemeBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnportalthemeBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnportaltheme_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["portaltheme"]; ok && val != nil {
		data.Portaltheme = types.StringValue(val.(string))
	} else {
		data.Portaltheme = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Portaltheme.ValueString())

	return data
}
