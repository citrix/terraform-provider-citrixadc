package bridgegroup_nsip_binding

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

// BridgegroupNsipBindingResourceModel describes the resource data model.
type BridgegroupNsipBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Bridgegroupid types.Int64  `tfsdk:"bridgegroup_id"`
	Ipaddress     types.String `tfsdk:"ipaddress"`
	Netmask       types.String `tfsdk:"netmask"`
	Ownergroup    types.String `tfsdk:"ownergroup"`
	Td            types.Int64  `tfsdk:"td"`
}

func (r *BridgegroupNsipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the bridgegroup_nsip_binding resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address assigned to the  bridge group.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The network mask for the subnet defined for the bridge group.",
			},
			"ownergroup": schema.StringAttribute{
				// Not echoed back by the NITRO GET response, so Computed would
				// leave it perpetually unknown after apply. Optional-only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				// Not echoed back by the NITRO GET response, so Computed would
				// leave it perpetually unknown after apply. Optional-only.
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func bridgegroup_nsip_bindingGetThePayloadFromthePlan(ctx context.Context, data *BridgegroupNsipBindingResourceModel) network.Bridgegroupnsipbinding {
	tflog.Debug(ctx, "In bridgegroup_nsip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// The bridge group integer key maps to the NITRO "id" field.
	bridgegroup_nsip_binding := network.Bridgegroupnsipbinding{}
	if !data.Bridgegroupid.IsNull() && !data.Bridgegroupid.IsUnknown() {
		bridgegroup_nsip_binding.Id = utils.IntPtr(int(data.Bridgegroupid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		bridgegroup_nsip_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		bridgegroup_nsip_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		bridgegroup_nsip_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		bridgegroup_nsip_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return bridgegroup_nsip_binding
}

// bridgegroupNsipBindingComputeId builds the composite ID in the new
// key:urlEncode(value) format, using the legacy attribute order
// (bridgegroup_id,ipaddress) recorded in resource_id_mapping.json so that
// ParseIdString round-trips both new and legacy SDK v2 IDs.
func bridgegroupNsipBindingComputeId(data *BridgegroupNsipBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%d", data.Bridgegroupid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(data.Ipaddress.ValueString())))
	return strings.Join(idParts, ",")
}

// bridgegroup_nsip_bindingSetAttrFromGet maps the GET response onto the model
// while preserving the resource's prior ID (set once in Create).
func bridgegroup_nsip_bindingSetAttrFromGet(ctx context.Context, data *BridgegroupNsipBindingResourceModel, getResponseData map[string]interface{}) *BridgegroupNsipBindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_nsip_bindingSetAttrFromGet Function")

	// Convert API response to model. The NITRO "id" field is the bridge group key.
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroupid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	}

	return data
}

// bridgegroup_nsip_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response and (re)computes the synthetic ID, since the
// datasource never runs Create.
func bridgegroup_nsip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *BridgegroupNsipBindingResourceModel, getResponseData map[string]interface{}) *BridgegroupNsipBindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_nsip_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroupid = types.Int64Value(intVal)
		}
	} else {
		data.Bridgegroupid = types.Int64Null()
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
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	data.Id = types.StringValue(bridgegroupNsipBindingComputeId(data))

	return data
}
