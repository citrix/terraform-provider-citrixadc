package vpnvserver_staserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverStaserverBindingResourceModel describes the resource data model.
type VpnvserverStaserverBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Staaddresstype types.String `tfsdk:"staaddresstype"`
	Staserver      types.String `tfsdk:"staserver"`
}

func (r *VpnvserverStaserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_staserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},
		},
	}
}

func vpnvserver_staserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverStaserverBindingResourceModel) vpn.Vpnvserverstaserverbinding {
	tflog.Debug(ctx, "In vpnvserver_staserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_staserver_binding := vpn.Vpnvserverstaserverbinding{}
	if !data.Name.IsNull() {
		vpnvserver_staserver_binding.Name = data.Name.ValueString()
	}
	if !data.Staaddresstype.IsNull() {
		vpnvserver_staserver_binding.Staaddresstype = data.Staaddresstype.ValueString()
	}
	if !data.Staserver.IsNull() {
		vpnvserver_staserver_binding.Staserver = data.Staserver.ValueString()
	}

	return vpnvserver_staserver_binding
}

func vpnvserver_staserver_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverStaserverBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverStaserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_staserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["staaddresstype"]; ok && val != nil {
		data.Staaddresstype = types.StringValue(val.(string))
	} else {
		data.Staaddresstype = types.StringNull()
	}
	if val, ok := getResponseData["staserver"]; ok && val != nil {
		data.Staserver = types.StringValue(val.(string))
	} else {
		data.Staserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("staserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Staserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
