package netbridge_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NetbridgeNsip6BindingResourceModel describes the resource data model.
type NetbridgeNsip6BindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *NetbridgeNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge_nsip6_binding resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The subnet that is extended by this network bridge.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the network bridge.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The network mask for the subnet.",
			},
		},
	}
}

func netbridge_nsip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *NetbridgeNsip6BindingResourceModel) network.Netbridgensip6binding {
	tflog.Debug(ctx, "In netbridge_nsip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	netbridge_nsip6_binding := network.Netbridgensip6binding{}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		netbridge_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		netbridge_nsip6_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		netbridge_nsip6_binding.Netmask = data.Netmask.ValueString()
	}

	return netbridge_nsip6_binding
}

// netbridge_nsip6_bindingSetAttrFromGet is the resource-side setter. It adopts the
// echoed-back values for name/ipaddress and, for netmask (server-overridden/omitted
// for IPv6 bindings), only adopts the GET value when present so the configured/state
// value is preserved (avoids "inconsistent result after apply"). It never recomputes
// the ID — the ID is set exactly once in Create. See Pattern 7/13.
func netbridge_nsip6_bindingSetAttrFromGet(ctx context.Context, data *NetbridgeNsip6BindingResourceModel, getResponseData map[string]interface{}) *NetbridgeNsip6BindingResourceModel {
	tflog.Debug(ctx, "In netbridge_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	return data
}

// netbridge_nsip6_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and sets the datasource ID (the datasource has no Create).
// See Pattern 7 datasource split.
func netbridge_nsip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *NetbridgeNsip6BindingResourceModel, getResponseData map[string]interface{}) *NetbridgeNsip6BindingResourceModel {
	tflog.Debug(ctx, "In netbridge_nsip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
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

	// Set ID for the datasource: matches resource_id_mapping.json order (name,ipaddress).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
