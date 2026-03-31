package vpnvserver_sharefileserver_binding

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

// VpnvserverSharefileserverBindingResourceModel describes the resource data model.
type VpnvserverSharefileserverBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Sharefile types.String `tfsdk:"sharefile"`
}

func (r *VpnvserverSharefileserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_sharefileserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"sharefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configured ShareFile server in XenMobile deployment. Format IP:PORT / FQDN:PORT",
			},
		},
	}
}

func vpnvserver_sharefileserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverSharefileserverBindingResourceModel) vpn.Vpnvserversharefileserverbinding {
	tflog.Debug(ctx, "In vpnvserver_sharefileserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_sharefileserver_binding := vpn.Vpnvserversharefileserverbinding{}
	if !data.Name.IsNull() {
		vpnvserver_sharefileserver_binding.Name = data.Name.ValueString()
	}
	if !data.Sharefile.IsNull() {
		vpnvserver_sharefileserver_binding.Sharefile = data.Sharefile.ValueString()
	}

	return vpnvserver_sharefileserver_binding
}

func vpnvserver_sharefileserver_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverSharefileserverBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverSharefileserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_sharefileserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["sharefile"]; ok && val != nil {
		data.Sharefile = types.StringValue(val.(string))
	} else {
		data.Sharefile = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sharefile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sharefile.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
