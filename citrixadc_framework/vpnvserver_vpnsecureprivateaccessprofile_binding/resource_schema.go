package vpnvserver_vpnsecureprivateaccessprofile_binding

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

// VpnvserverVpnsecureprivateaccessprofileBindingResourceModel describes the resource data model.
type VpnvserverVpnsecureprivateaccessprofileBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Secureprivateaccessprofile types.String `tfsdk:"secureprivateaccessprofile"`
	Name                       types.String `tfsdk:"name"`
}

func (r *VpnvserverVpnsecureprivateaccessprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnsecureprivateaccessprofile_binding resource.",
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Secure Private Access profile bound to the vserver.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server.",
			},
		},
	}
}

func vpnvserver_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverVpnsecureprivateaccessprofileBindingResourceModel) vpn.Vpnvservervpnsecureprivateaccessprofilebinding {
	tflog.Debug(ctx, "In vpnvserver_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_vpnsecureprivateaccessprofile_binding := vpn.Vpnvservervpnsecureprivateaccessprofilebinding{}
	if !data.Secureprivateaccessprofile.IsNull() && !data.Secureprivateaccessprofile.IsUnknown() {
		vpnvserver_vpnsecureprivateaccessprofile_binding.Secureprivateaccessprofile = data.Secureprivateaccessprofile.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_vpnsecureprivateaccessprofile_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_vpnsecureprivateaccessprofile_binding
}

func vpnvserver_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnsecureprivateaccessprofileBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnsecureprivateaccessprofileBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnsecureprivateaccessprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["secureprivateaccessprofile"]; ok && val != nil {
		data.Secureprivateaccessprofile = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessprofile = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
