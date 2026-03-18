package vpnvserver_vpneula_binding

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

// VpnvserverVpneulaBindingResourceModel describes the resource data model.
type VpnvserverVpneulaBindingResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Eula types.String `tfsdk:"eula"`
	Name types.String `tfsdk:"name"`
}

func (r *VpnvserverVpneulaBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpneula_binding resource.",
			},
			"eula": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the EULA bound to VPN vserver",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}

func vpnvserver_vpneula_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverVpneulaBindingResourceModel) vpn.Vpnvservervpneulabinding {
	tflog.Debug(ctx, "In vpnvserver_vpneula_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_vpneula_binding := vpn.Vpnvservervpneulabinding{}
	if !data.Eula.IsNull() {
		vpnvserver_vpneula_binding.Eula = data.Eula.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_vpneula_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_vpneula_binding
}

func vpnvserver_vpneula_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpneulaBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpneulaBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpneula_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["eula"]; ok && val != nil {
		data.Eula = types.StringValue(val.(string))
	} else {
		data.Eula = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("eula:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Eula.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
