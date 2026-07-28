package vlan_linkset_binding

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

// VlanLinksetBindingResourceModel describes the resource data model.
//
// NOTE: there is intentionally NO `linkset` attribute. Despite the resource name,
// the NITRO endpoint vlan_linkset_binding (Go struct network.Vlanlinksetbinding)
// binds an INTERFACE (ifnum), with optional tagged/ownergroup, to a VLAN (id).
// The NITRO doc, the vendored struct, and tfdata all expose only: id, ifnum,
// tagged, ownergroup. The user-facing VLAN id is mapped to `vlanid` (Int64) to
// avoid a name collision with the synthetic Terraform `id`, matching the sibling
// vlan_channel_binding / vlan_interface_binding resources.
type VlanLinksetBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Vlanid     types.Int64  `tfsdk:"vlanid"`
	Ifnum      types.String `tfsdk:"ifnum"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Tagged     types.Bool   `tfsdk:"tagged"`
}

func (r *VlanLinksetBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vlan_linkset_binding resource.",
			},
			"vlanid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the virtual LAN ID.",
			},
			"ifnum": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The interface to be bound to the VLAN, specified in slot/port notation (for example, 1/3).",
			},
			"ownergroup": schema.StringAttribute{
				// NOTE: ownergroup is a CLUSTER-only attribute. It must NOT be sent on a
				// standalone appliance: doing so triggers NITRO errorcode 1093
				// "Argument pre-requisite missing [ownerGroup, IPAddress]". Previously this
				// attribute carried Computed + Default("DEFAULT_NG"), which forced the value
				// into every create payload and made the resource unusable on standalone
				// ADCs. It is now purely Optional so it is only sent when the user sets it.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"tagged": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Make the interface an 802.1q tagged interface. Packets sent on this interface on this VLAN have an additional 4-byte 802.1q tag, which identifies the VLAN. To use 802.1q tagging, you must also configure the switch connected to the appliance's interfaces.",
			},
		},
	}
}

func vlan_linkset_bindingGetThePayloadFromthePlan(ctx context.Context, data *VlanLinksetBindingResourceModel) network.Vlanlinksetbinding {
	tflog.Debug(ctx, "In vlan_linkset_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vlan_linkset_binding := network.Vlanlinksetbinding{}
	if !data.Vlanid.IsNull() && !data.Vlanid.IsUnknown() {
		vlan_linkset_binding.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		vlan_linkset_binding.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		vlan_linkset_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Tagged.IsNull() && !data.Tagged.IsUnknown() {
		vlan_linkset_binding.Tagged = data.Tagged.ValueBool()
	}

	return vlan_linkset_binding
}

// vlan_linkset_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (vlanid, ifnum are
// RequiresReplace identity attrs) and does NOT recompute the ID, which is set
// exactly once in Create (Pattern 6/7).
func vlan_linkset_bindingSetAttrFromGet(ctx context.Context, data *VlanLinksetBindingResourceModel, getResponseData map[string]interface{}) *VlanLinksetBindingResourceModel {
	tflog.Debug(ctx, "In vlan_linkset_bindingSetAttrFromGet Function")

	// Convert API response to model. Only adopt server-echoed values; preserve
	// configured identity values otherwise (RequiresReplace attrs).
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

// vlan_linkset_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create
// to seed those values (Pattern 7 datasource split).
func vlan_linkset_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VlanLinksetBindingResourceModel, getResponseData map[string]interface{}) *VlanLinksetBindingResourceModel {
	tflog.Debug(ctx, "In vlan_linkset_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
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

	// Set ID for the datasource
	// Composite key: vlanid,ifnum (key:UrlEncode(value) pairs)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
