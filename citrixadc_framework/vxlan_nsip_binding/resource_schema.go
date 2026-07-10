package vxlan_nsip_binding

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

// VxlanNsipBindingResourceModel describes the resource data model.
type VxlanNsipBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Vxlanid   types.Int64  `tfsdk:"vxlanid"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *VxlanNsipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlan_nsip_binding resource.",
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

func vxlan_nsip_bindingGetThePayloadFromthePlan(ctx context.Context, data *VxlanNsipBindingResourceModel) network.Vxlannsipbinding {
	tflog.Debug(ctx, "In vxlan_nsip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NITRO carries the vxlanid in the struct field "Id" (json:"id").
	vxlan_nsip_binding := network.Vxlannsipbinding{}
	if !data.Vxlanid.IsNull() && !data.Vxlanid.IsUnknown() {
		vxlan_nsip_binding.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		vxlan_nsip_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		vxlan_nsip_binding.Netmask = data.Netmask.ValueString()
	}

	return vxlan_nsip_binding
}

// vxlan_nsip_bindingComposeId builds the composite resource ID matching the legacy
// SDK v2 attribute order (vxlanid,ipaddress) in the new key:value format.
func vxlan_nsip_bindingComposeId(data *VxlanNsipBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vxlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(data.Ipaddress.ValueString())))
	return strings.Join(idParts, ",")
}

// vxlan_nsip_bindingSetAttrFromGet preserves the configured/state values for the
// resource. The NITRO GET response normalizes/overrides server-side fields, so the
// resource keeps its planned values (Pattern 7).
func vxlan_nsip_bindingSetAttrFromGet(ctx context.Context, data *VxlanNsipBindingResourceModel, getResponseData map[string]interface{}) *VxlanNsipBindingResourceModel {
	tflog.Debug(ctx, "In vxlan_nsip_bindingSetAttrFromGet Function")

	// vxlanid is the parent key (returned as "id" by NITRO). Only adopt the GET value
	// when the model does not already carry it (e.g. import).
	if data.Vxlanid.IsNull() || data.Vxlanid.IsUnknown() {
		if val, ok := getResponseData["id"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Vxlanid = types.Int64Value(intVal)
			}
		}
	}
	if data.Ipaddress.IsNull() || data.Ipaddress.IsUnknown() {
		if val, ok := getResponseData["ipaddress"]; ok && val != nil {
			data.Ipaddress = types.StringValue(val.(string))
		}
	}
	// netmask is server-overridable; adopt the value from GET to reflect actual state.
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	// Set the composite ID for the resource.
	data.Id = types.StringValue(vxlan_nsip_bindingComposeId(data))

	return data
}

// vxlan_nsip_bindingSetAttrFromGetForDatasource faithfully copies every field from the
// GET response (the datasource has no prior plan/state to preserve) and sets the ID.
func vxlan_nsip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VxlanNsipBindingResourceModel, getResponseData map[string]interface{}) *VxlanNsipBindingResourceModel {
	tflog.Debug(ctx, "In vxlan_nsip_bindingSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue(vxlan_nsip_bindingComposeId(data))

	return data
}
