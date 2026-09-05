package netbridge_nsip_binding

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

// NetbridgeNsipBindingResourceModel describes the resource data model.
type NetbridgeNsipBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *NetbridgeNsipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge_nsip_binding resource.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The network mask for the subnet.",
			},
		},
	}
}

func netbridge_nsip_bindingGetThePayloadFromthePlan(ctx context.Context, data *NetbridgeNsipBindingResourceModel) network.Netbridgensipbinding {
	tflog.Debug(ctx, "In netbridge_nsip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	netbridge_nsip_binding := network.Netbridgensipbinding{}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		netbridge_nsip_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		netbridge_nsip_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		netbridge_nsip_binding.Netmask = data.Netmask.ValueString()
	}

	return netbridge_nsip_binding
}

// netbridge_nsip_bindingSetAttrFromGet is the resource-side state setter.
// The NITRO GET response for this binding does NOT echo back netmask (the SDK v2
// resource explicitly skipped d.Set("netmask", ...) for the same reason), so we
// preserve the prior plan/state value for netmask instead of nulling it (Pattern 7).
// We also do NOT recompute data.Id here — the ID is composed once in Create and
// must not be rebuilt from a GET response that lacks netmask (Pattern 6).
func netbridge_nsip_bindingSetAttrFromGet(ctx context.Context, data *NetbridgeNsipBindingResourceModel, getResponseData map[string]interface{}) *NetbridgeNsipBindingResourceModel {
	tflog.Debug(ctx, "In netbridge_nsip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// netmask is not returned by the GET response; preserve the existing plan/state value.
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// netbridge_nsip_bindingSetAttrFromGetForDatasource is the datasource-side setter.
// A datasource has no prior plan/state to preserve, so it faithfully copies the
// fields from the GET response (netmask is supplied from config in the datasource
// Read) and composes the ID itself, since the datasource never calls Create (Pattern 7).
func netbridge_nsip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *NetbridgeNsipBindingResourceModel, getResponseData map[string]interface{}) *NetbridgeNsipBindingResourceModel {
	tflog.Debug(ctx, "In netbridge_nsip_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// netmask is not returned by GET; the datasource keeps the config-provided value.

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
