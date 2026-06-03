package vpnvserver_vpnepaprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverVpnepaprofileBindingResourceModel describes the resource data model.
type VpnvserverVpnepaprofileBindingResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Epaprofile         types.String `tfsdk:"epaprofile"`
	Epaprofileoptional types.Bool   `tfsdk:"epaprofileoptional"`
	Name               types.String `tfsdk:"name"`
}

func (r *VpnvserverVpnepaprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnepaprofile_binding resource.",
			},
			"epaprofile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Advanced EPA profile to bind",
			},
			"epaprofileoptional": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Mark the EPA profile optional for preauthentication EPA profile. User would be shown a logon page even if the EPA profile fails to evaluate.",
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

func vpnvserver_vpnepaprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverVpnepaprofileBindingResourceModel) vpn.Vpnvservervpnepaprofilebinding {
	tflog.Debug(ctx, "In vpnvserver_vpnepaprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_vpnepaprofile_binding := vpn.Vpnvservervpnepaprofilebinding{}
	if !data.Epaprofile.IsNull() && !data.Epaprofile.IsUnknown() {
		vpnvserver_vpnepaprofile_binding.Epaprofile = data.Epaprofile.ValueString()
	}
	if !data.Epaprofileoptional.IsNull() && !data.Epaprofileoptional.IsUnknown() {
		vpnvserver_vpnepaprofile_binding.Epaprofileoptional = data.Epaprofileoptional.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_vpnepaprofile_binding.Name = data.Name.ValueString()
	}

	return vpnvserver_vpnepaprofile_binding
}

func vpnvserver_vpnepaprofile_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnepaprofileBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnepaprofileBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnepaprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["epaprofile"]; ok && val != nil {
		data.Epaprofile = types.StringValue(val.(string))
	} else {
		data.Epaprofile = types.StringNull()
	}
	if val, ok := getResponseData["epaprofileoptional"]; ok && val != nil {
		data.Epaprofileoptional = types.BoolValue(val.(bool))
	} else {
		data.Epaprofileoptional = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("epaprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Epaprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
