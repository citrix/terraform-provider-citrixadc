package vpnvserver_vpnintranetapplication_binding

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

// VpnvserverVpnintranetapplicationBindingResourceModel describes the resource data model.
type VpnvserverVpnintranetapplicationBindingResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Intranetapplication types.String `tfsdk:"intranetapplication"`
	Name                types.String `tfsdk:"name"`
}

func (r *VpnvserverVpnintranetapplicationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnintranetapplication_binding resource.",
			},
			"intranetapplication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet VPN application.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}

func vpnvserver_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverVpnintranetapplicationBindingResourceModel) vpn.Vpnvservervpnintranetapplicationbinding {
	tflog.Debug(ctx, "In vpnvserver_vpnintranetapplication_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_vpnintranetapplication_binding := vpn.Vpnvservervpnintranetapplicationbinding{}
	if !data.Intranetapplication.IsNull() {
		vpnvserver_vpnintranetapplication_binding.Intranetapplication = data.Intranetapplication.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_vpnintranetapplication_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_vpnintranetapplication_binding
}

func vpnvserver_vpnintranetapplication_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnintranetapplicationBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnintranetapplicationBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnintranetapplication_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["intranetapplication"]; ok && val != nil {
		data.Intranetapplication = types.StringValue(val.(string))
	} else {
		data.Intranetapplication = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetapplication:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Intranetapplication.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
