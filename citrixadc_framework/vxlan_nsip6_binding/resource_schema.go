package vxlan_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VxlanNsip6BindingResourceModel describes the resource data model.
type VxlanNsip6BindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Vxlanid   types.Int64  `tfsdk:"vxlanid"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *VxlanNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlan_nsip6_binding resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vxlanid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address assigned to the VXLAN.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the network address defined for this VXLAN.",
			},
		},
	}
}

func vxlan_nsip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *VxlanNsip6BindingResourceModel) network.Vxlannsip6binding {
	tflog.Debug(ctx, "In vxlan_nsip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vxlan_nsip6_binding := network.Vxlannsip6binding{}
	if !data.Vxlanid.IsNull() && !data.Vxlanid.IsUnknown() {
		vxlan_nsip6_binding.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		vxlan_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		vxlan_nsip6_binding.Netmask = data.Netmask.ValueString()
	}

	return vxlan_nsip6_binding
}

// vxlan_nsip6_bindingSetAttrFromGet is the resource-side state setter. It
// preserves the existing composite ID (set once in Create) and maps the
// server-returned fields back into the model. The integer VXLAN key is
// returned by NITRO under "id" and stored under the user-facing "vxlanid".
func vxlan_nsip6_bindingSetAttrFromGet(ctx context.Context, data *VxlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VxlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vxlan_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	return data
}

// vxlan_nsip6_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the datasource ID (the datasource has no
// Create to set it).
func vxlan_nsip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VxlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VxlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vxlan_nsip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vxlanid = types.Int64Null()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	data.Id = types.StringValue(vxlan_nsip6_bindingComposeId(data))

	return data
}

// vxlan_nsip6_bindingComposeId builds the new key:value composite ID using the
// legacy attribute order (vxlanid,ipaddress) so it round-trips with
// resource_id_mapping.json.
func vxlan_nsip6_bindingComposeId(data *VxlanNsip6BindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vxlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(data.Ipaddress.ValueString())))
	return strings.Join(idParts, ",")
}
