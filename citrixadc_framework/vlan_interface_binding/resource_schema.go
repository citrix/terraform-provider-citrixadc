package vlan_interface_binding

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

// VlanInterfaceBindingResourceModel describes the resource data model.
type VlanInterfaceBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Vlanid     types.Int64  `tfsdk:"vlanid"`
	Ifnum      types.String `tfsdk:"ifnum"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Tagged     types.Bool   `tfsdk:"tagged"`
}

func (r *VlanInterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vlan_interface_binding resource.",
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
				// Optional only (not Computed): the NITRO GET for this binding does
				// not echo ownergroup on a standalone ADC, and sending it
				// unconditionally triggers errorcode 1093 (requires IPAddress). When
				// omitted it stays null, mirroring the SDK v2 behavior (the value was
				// only sent when the user configured it).
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

func vlan_interface_bindingGetThePayloadFromthePlan(ctx context.Context, data *VlanInterfaceBindingResourceModel) network.Vlaninterfacebinding {
	tflog.Debug(ctx, "In vlan_interface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model. The user-facing "vlanid" attribute
	// maps to the NITRO struct field Id (json "id").
	vlan_interface_binding := network.Vlaninterfacebinding{}
	if !data.Vlanid.IsNull() && !data.Vlanid.IsUnknown() {
		vlan_interface_binding.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		vlan_interface_binding.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		vlan_interface_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Tagged.IsNull() && !data.Tagged.IsUnknown() {
		vlan_interface_binding.Tagged = data.Tagged.ValueBool()
	}

	return vlan_interface_binding
}

// vlanInterfaceBindingComposeId builds the composite ID matching the legacy
// SDK v2 attribute order (resource_id_mapping.json: "vlanid,ifnum").
func vlanInterfaceBindingComposeId(data *VlanInterfaceBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	return strings.Join(idParts, ",")
}

// vlan_interface_bindingSetAttrFromGet is the resource-side setter. It preserves
// the user-configured key attributes (vlanid, ifnum) while adopting the
// server-defaulted fields (ownergroup, tagged) from the GET response.
func vlan_interface_bindingSetAttrFromGet(ctx context.Context, data *VlanInterfaceBindingResourceModel, getResponseData map[string]interface{}) *VlanInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In vlan_interface_bindingSetAttrFromGet Function")

	// vlanid (NITRO field "id") and ifnum are identity attributes (RequiresReplace);
	// adopt the GET value (it equals the configured value).
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	}
	// ownergroup and tagged are server-defaulted; adopt the GET value.
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["tagged"]; ok && val != nil {
		data.Tagged = types.BoolValue(val.(bool))
	}

	// Compose the composite ID (vlanid:..,ifnum:..)
	data.Id = types.StringValue(vlanInterfaceBindingComposeId(data))

	return data
}

// vlan_interface_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response (datasources have no prior state to preserve) and sets the ID.
func vlan_interface_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VlanInterfaceBindingResourceModel, getResponseData map[string]interface{}) *VlanInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In vlan_interface_bindingSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue(vlanInterfaceBindingComposeId(data))

	return data
}
