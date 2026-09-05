package vpnvserver_intranetip_binding

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

// VpnvserverIntranetipBindingResourceModel describes the resource data model.
type VpnvserverIntranetipBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Intranetip types.String `tfsdk:"intranetip"`
	Name       types.String `tfsdk:"name"`
	Netmask    types.String `tfsdk:"netmask"`
}

func (r *VpnvserverIntranetipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_intranetip_binding resource.",
			},
			"intranetip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The netmask of the intranet IP address or range.",
			},
		},
	}
}

func vpnvserver_intranetip_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverIntranetipBindingResourceModel) vpn.Vpnvserverintranetipbinding {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnvserver_intranetip_binding := vpn.Vpnvserverintranetipbinding{}
	if !data.Intranetip.IsNull() && !data.Intranetip.IsUnknown() {
		vpnvserver_intranetip_binding.Intranetip = data.Intranetip.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver_intranetip_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		vpnvserver_intranetip_binding.Netmask = data.Netmask.ValueString()
	}

	return vpnvserver_intranetip_binding
}

// vpnvserver_intranetip_bindingSetAttrFromGet is the resource-side state setter.
// It copies the GET response into the model but does NOT recompute data.Id —
// the ID is set exactly once in Create (Pattern 6). netmask is server-overridable
// and not always echoed (Pattern 7/13): preserve the existing plan/state value when
// the GET response does not carry it, instead of nulling it.
func vpnvserver_intranetip_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// vpnvserver_intranetip_bindingSetAttrFromGetForDatasource is the datasource-side
// setter. The datasource has no prior plan/state, so it faithfully copies every
// field from the GET response and sets data.Id (Pattern 7 datasource split).
func vpnvserver_intranetip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	} else {
		data.Intranetip = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the datasource using the legacy attribute order (name, intranetip)
	// matching resource_id_mapping.json.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
