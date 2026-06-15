package vlan_channel_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VlanChannelBindingResourceModel describes the resource data model.
type VlanChannelBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Vlanid     types.Int64  `tfsdk:"vlanid"`
	Ifnum      types.String `tfsdk:"ifnum"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Tagged     types.Bool   `tfsdk:"tagged"`
}

func (r *VlanChannelBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vlan_channel_binding resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vlanid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the virtual LAN ID.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).",
			},
			"ownergroup": schema.StringAttribute{
				// Optional only (no Computed): the binding GET response does not
				// echo ownergroup back, so a Computed flag would leave the value
				// perpetually unknown after apply.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"tagged": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.",
			},
		},
	}
}

func vlan_channel_bindingGetThePayloadFromthePlan(ctx context.Context, data *VlanChannelBindingResourceModel) network.Vlanchannelbinding {
	tflog.Debug(ctx, "In vlan_channel_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vlan_channel_binding := network.Vlanchannelbinding{}
	if !data.Vlanid.IsNull() && !data.Vlanid.IsUnknown() {
		vlan_channel_binding.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		vlan_channel_binding.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		vlan_channel_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Tagged.IsNull() && !data.Tagged.IsUnknown() {
		vlan_channel_binding.Tagged = data.Tagged.ValueBool()
	}

	return vlan_channel_binding
}

// vlan_channel_bindingComposeId builds the composite resource ID using the
// legacy SDK v2 attribute order (vlanid,ifnum) in the new key:value form.
func vlan_channel_bindingComposeId(data *VlanChannelBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(data.Ifnum.ValueString())))
	return strings.Join(idParts, ",")
}

// vlan_channel_bindingSetAttrFromGet is the resource-side state setter. It
// preserves the synthetic Id (set once in Create) and refreshes the
// server-managed fields from the GET response.
func vlan_channel_bindingSetAttrFromGet(ctx context.Context, data *VlanChannelBindingResourceModel, getResponseData map[string]interface{}) *VlanChannelBindingResourceModel {
	tflog.Debug(ctx, "In vlan_channel_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.BoolValue(val.(bool))
	}

	return data
}

// vlan_channel_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and sets the synthetic Id, for use by the datasource
// which has no Create to seed the ID.
func vlan_channel_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VlanChannelBindingResourceModel, getResponseData map[string]interface{}) *VlanChannelBindingResourceModel {
	tflog.Debug(ctx, "In vlan_channel_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vlanid = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.BoolValue(val.(bool))
	} else {
		data.Tagged = types.BoolNull()
	}

	data.Id = types.StringValue(vlan_channel_bindingComposeId(data))

	return data
}
