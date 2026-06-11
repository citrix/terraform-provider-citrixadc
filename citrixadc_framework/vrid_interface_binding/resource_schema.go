package vrid_interface_binding

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

// VridInterfaceBindingResourceModel describes the RESOURCE data model.
//
// NOTE: the NITRO attribute "id" (the integer VRID) is exposed as "vrid_id" in
// the Terraform schema to avoid colliding with the framework's synthetic string
// "id" attribute. The JSON wire field stays "id".
//
// flags and vlan are NITRO read-only output fields; they are NOT writable inputs
// to the bind (PUT) and are exposed only on the DATASOURCE model below.
type VridInterfaceBindingResourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"`
	Ifnum  types.String `tfsdk:"ifnum"`
}

// VridInterfaceBindingDataSourceModel describes the DATASOURCE data model. It adds
// the read-only output fields (flags, vlan) returned by the GET endpoint.
type VridInterfaceBindingDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Flags  types.Int64  `tfsdk:"flags"`
	Vlan   types.Int64  `tfsdk:"vlan"`
}

func (r *VridInterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid_interface_binding resource.",
			},
			"vrid_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interfaces to bind to the VMAC, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},
		},
	}
}

func vrid_interface_bindingGetThePayloadFromthePlan(ctx context.Context, data *VridInterfaceBindingResourceModel) network.Vridinterfacebinding {
	tflog.Debug(ctx, "In vrid_interface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only id and ifnum are write fields;
	// flags and vlan are read-only and never sent.
	vrid_interface_binding := network.Vridinterfacebinding{}
	if !data.VridId.IsNull() && !data.VridId.IsUnknown() {
		vrid_interface_binding.Id = utils.IntPtr(int(data.VridId.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		vrid_interface_binding.Ifnum = data.Ifnum.ValueString()
	}

	return vrid_interface_binding
}

// vrid_interface_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (vrid_id, ifnum are both
// RequiresReplace) and does NOT recompute the ID, which is set exactly once in Create.
func vrid_interface_bindingSetAttrFromGet(ctx context.Context, data *VridInterfaceBindingResourceModel, getResponseData map[string]interface{}) *VridInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In vrid_interface_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	}

	return data
}

// vrid_interface_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response (including the read-only flags and vlan) and composes the
// ID, since the datasource has no Create to seed those values.
func vrid_interface_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VridInterfaceBindingDataSourceModel, getResponseData map[string]interface{}) *VridInterfaceBindingDataSourceModel {
	tflog.Debug(ctx, "In vrid_interface_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	} else {
		data.VridId = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		// NITRO may echo ifnum as a scalar string or (on some firmware) a list.
		switch t := val.(type) {
		case string:
			data.Ifnum = types.StringValue(t)
		case []interface{}:
			if len(t) > 0 {
				if s, ok := t[0].(string); ok {
					data.Ifnum = types.StringValue(s)
				}
			}
		}
	}
	// NOTE: the firmware does NOT echo "ifnum" for vrid_interface_binding rows. When
	// it is absent we deliberately RETAIN the config-supplied ifnum (the datasource's
	// lookup key) instead of nulling it, so the composite ID and the ifnum output
	// remain correct.
	if val, ok := getResponseData["flags"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Flags = types.Int64Value(intVal)
		}
	} else {
		data.Flags = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the datasource. Composite key: vrid_id (NITRO id), ifnum.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
