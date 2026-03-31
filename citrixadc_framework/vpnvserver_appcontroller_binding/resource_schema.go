package vpnvserver_appcontroller_binding

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

// VpnvserverAppcontrollerBindingResourceModel describes the resource data model.
type VpnvserverAppcontrollerBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Appcontroller types.String `tfsdk:"appcontroller"`
	Name          types.String `tfsdk:"name"`
}

func (r *VpnvserverAppcontrollerBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_appcontroller_binding resource.",
			},
			"appcontroller": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured App Controller server in XenMobile deployment.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}

func vpnvserver_appcontroller_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverAppcontrollerBindingResourceModel) vpn.Vpnvserverappcontrollerbinding {
	tflog.Debug(ctx, "In vpnvserver_appcontroller_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_appcontroller_binding := vpn.Vpnvserverappcontrollerbinding{}
	if !data.Appcontroller.IsNull() {
		vpnvserver_appcontroller_binding.Appcontroller = data.Appcontroller.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_appcontroller_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_appcontroller_binding
}

func vpnvserver_appcontroller_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAppcontrollerBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAppcontrollerBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_appcontroller_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appcontroller"]; ok && val != nil {
		data.Appcontroller = types.StringValue(val.(string))
	} else {
		data.Appcontroller = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("appcontroller:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Appcontroller.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
