package vpnvserver_analyticsprofile_binding

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

// VpnvserverAnalyticsprofileBindingResourceModel describes the resource data model.
type VpnvserverAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
	Name             types.String `tfsdk:"name"`
}

func (r *VpnvserverAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the analytics profile bound to the VPN Vserver",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
		},
	}
}

func vpnvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverAnalyticsprofileBindingResourceModel) vpn.Vpnvserveranalyticsprofilebinding {
	tflog.Debug(ctx, "In vpnvserver_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_analyticsprofile_binding := vpn.Vpnvserveranalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		vpnvserver_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_analyticsprofile_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_analyticsprofile_binding
}

func vpnvserver_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_analyticsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else {
		data.Analyticsprofile = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("analyticsprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
