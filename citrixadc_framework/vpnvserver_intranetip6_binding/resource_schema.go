package vpnvserver_intranetip6_binding

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

// VpnvserverIntranetip6BindingResourceModel describes the resource data model.
type VpnvserverIntranetip6BindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Intranetip6 types.String `tfsdk:"intranetip6"`
	Name        types.String `tfsdk:"name"`
	Numaddr     types.Int64  `tfsdk:"numaddr"`
}

func (r *VpnvserverIntranetip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_intranetip6_binding resource.",
			},
			"intranetip6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"numaddr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of ipv6 addresses",
			},
		},
	}
}

func vpnvserver_intranetip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverIntranetip6BindingResourceModel) vpn.Vpnvserverintranetip6binding {
	tflog.Debug(ctx, "In vpnvserver_intranetip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_intranetip6_binding := vpn.Vpnvserverintranetip6binding{}
	if !data.Intranetip6.IsNull() {
		vpnvserver_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_intranetip6_binding.Name = data.Name.ValueString()
	}
	if !data.Numaddr.IsNull() {
		vpnvserver_intranetip6_binding.Numaddr = utils.IntPtr(int(data.Numaddr.ValueInt64()))
	}

	return vpnvserver_intranetip6_binding
}

func vpnvserver_intranetip6_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *VpnvserverIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_intranetip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	} else {
		data.Intranetip6 = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	} else {
		data.Numaddr = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip6.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
