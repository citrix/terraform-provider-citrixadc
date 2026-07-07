package bridgegroup_nsip6_binding

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

// BridgegroupNsip6BindingResourceModel describes the resource data model.
type BridgegroupNsip6BindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	BridgegroupId types.Int64  `tfsdk:"bridgegroup_id"`
	Ipaddress     types.String `tfsdk:"ipaddress"`
	Netmask       types.String `tfsdk:"netmask"`
	Ownergroup    types.String `tfsdk:"ownergroup"`
	Td            types.Int64  `tfsdk:"td"`
}

func (r *BridgegroupNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the bridgegroup_nsip6_binding resource.",
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
			// netmask/ownergroup/td are NOT echoed by the binding GET response, so
			// they cannot be Computed (Terraform would see an unresolved unknown after
			// apply). Keep them Optional-only — the configured value is preserved.
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A subnet mask associated with the network address.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func bridgegroup_nsip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *BridgegroupNsip6BindingResourceModel) network.Bridgegroupnsip6binding {
	tflog.Debug(ctx, "In bridgegroup_nsip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	bridgegroup_nsip6_binding := network.Bridgegroupnsip6binding{}
	if !data.BridgegroupId.IsNull() && !data.BridgegroupId.IsUnknown() {
		bridgegroup_nsip6_binding.Id = utils.IntPtr(int(data.BridgegroupId.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		bridgegroup_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		bridgegroup_nsip6_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		bridgegroup_nsip6_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		bridgegroup_nsip6_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return bridgegroup_nsip6_binding
}

// bridgegroup_nsip6_bindingComposeId builds the composite resource ID using the
// legacy SDK v2 attribute order (bridgegroup_id, ipaddress) in the new key:value form.
func bridgegroup_nsip6_bindingComposeId(data *BridgegroupNsip6BindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BridgegroupId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(data.Ipaddress.ValueString())))
	return strings.Join(idParts, ",")
}

// bridgegroup_nsip6_bindingSetAttrFromGet populates the resource model from a GET
// response while preserving the synthetic composite ID (set once in Create).
func bridgegroup_nsip6_bindingSetAttrFromGet(ctx context.Context, data *BridgegroupNsip6BindingResourceModel, getResponseData map[string]interface{}) *BridgegroupNsip6BindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
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

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(bridgegroup_nsip6_bindingComposeId(data))

	return data
}

// bridgegroup_nsip6_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response and sets the datasource ID (datasources have no Create).
func bridgegroup_nsip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *BridgegroupNsip6BindingResourceModel, getResponseData map[string]interface{}) *BridgegroupNsip6BindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_nsip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
		}
	} else {
		data.BridgegroupId = types.Int64Null()
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

	// Set ID for the datasource
	data.Id = types.StringValue(bridgegroup_nsip6_bindingComposeId(data))

	return data
}
