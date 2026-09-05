package vrid_trackinterface_binding

import (
	"context"

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

// VridTrackinterfaceBindingResourceModel describes the RESOURCE data model.
//
// NOTE: the NITRO attribute "id" (the integer VRID) is exposed as "vrid_id" in
// the Terraform schema to avoid colliding with the framework's synthetic string
// "id" attribute. The JSON wire field stays "id".
//
// flags is a NITRO read-only output field; it is NOT a writable input to the bind
// (PUT) and is exposed only on the DATASOURCE model below. (This binding has no
// vlan field, unlike vrid_channel_binding / vrid_interface_binding.)
type VridTrackinterfaceBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	VridId     types.Int64  `tfsdk:"vrid_id"`
	Trackifnum types.String `tfsdk:"trackifnum"`
}

func (r *VridTrackinterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid_trackinterface_binding resource.",
			},
			"vrid_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"trackifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interfaces which need to be tracked for this vrID.",
			},
		},
	}
}

func vrid_trackinterface_bindingGetThePayloadFromthePlan(ctx context.Context, data *VridTrackinterfaceBindingResourceModel) network.Vridtrackinterfacebinding {
	tflog.Debug(ctx, "In vrid_trackinterface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only id and trackifnum are write
	// fields; flags is read-only and never sent.
	vrid_trackinterface_binding := network.Vridtrackinterfacebinding{}
	if !data.VridId.IsNull() && !data.VridId.IsUnknown() {
		vrid_trackinterface_binding.Id = utils.IntPtr(int(data.VridId.ValueInt64()))
	}
	if !data.Trackifnum.IsNull() && !data.Trackifnum.IsUnknown() {
		vrid_trackinterface_binding.Trackifnum = data.Trackifnum.ValueString()
	}

	return vrid_trackinterface_binding
}

// vrid_trackinterface_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (vrid_id, trackifnum are both
// RequiresReplace) and does NOT recompute the ID, which is set exactly once in Create.
func vrid_trackinterface_bindingSetAttrFromGet(ctx context.Context, data *VridTrackinterfaceBindingResourceModel, getResponseData map[string]interface{}) *VridTrackinterfaceBindingResourceModel {
	tflog.Debug(ctx, "In vrid_trackinterface_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["trackifnum"]; ok && val != nil {
		data.Trackifnum = types.StringValue(val.(string))
	}

	return data
}
