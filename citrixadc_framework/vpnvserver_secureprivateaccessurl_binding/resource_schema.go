package vpnvserver_secureprivateaccessurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverSecureprivateaccessurlBindingResourceModel describes the resource data model.
type VpnvserverSecureprivateaccessurlBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`
	Secureprivateaccessurl types.String `tfsdk:"secureprivateaccessurl"`
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_secureprivateaccessurl_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server.",
			},
			"secureprivateaccessurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configured Secure Private Access URL",
			},
		},
	}
}

func vpnvserver_secureprivateaccessurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverSecureprivateaccessurlBindingResourceModel) vpn.Vpnvserversecureprivateaccessurlbinding {
	tflog.Debug(ctx, "In vpnvserver_secureprivateaccessurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_secureprivateaccessurl_binding := vpn.Vpnvserversecureprivateaccessurlbinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_secureprivateaccessurl_binding.Name = data.Name.ValueString()
	}
	if !data.Secureprivateaccessurl.IsNull() && !data.Secureprivateaccessurl.IsUnknown() {
		vpnvserver_secureprivateaccessurl_binding.Secureprivateaccessurl = data.Secureprivateaccessurl.ValueString()
	}

	return vpnvserver_secureprivateaccessurl_binding
}

func vpnvserver_secureprivateaccessurl_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverSecureprivateaccessurlBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverSecureprivateaccessurlBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_secureprivateaccessurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccessurl"]; ok && val != nil {
		data.Secureprivateaccessurl = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessurl = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
