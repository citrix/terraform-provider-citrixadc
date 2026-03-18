package vpnglobal_vpnurl_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalVpnurlBindingResourceModel describes the resource data model.
type VpnglobalVpnurlBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Urlname                types.String `tfsdk:"urlname"`
}

func (r *VpnglobalVpnurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnurl_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"urlname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet url.",
			},
		},
	}
}

func vpnglobal_vpnurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalVpnurlBindingResourceModel) vpn.Vpnglobalvpnurlbinding {
	tflog.Debug(ctx, "In vpnglobal_vpnurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_vpnurl_binding := vpn.Vpnglobalvpnurlbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_vpnurl_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Urlname.IsNull() {
		vpnglobal_vpnurl_binding.Urlname = data.Urlname.ValueString()
	}

	return vpnglobal_vpnurl_binding
}

func vpnglobal_vpnurl_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnurlBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnurlBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["urlname"]; ok && val != nil {
		data.Urlname = types.StringValue(val.(string))
	} else {
		data.Urlname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Urlname.ValueString())

	return data
}
