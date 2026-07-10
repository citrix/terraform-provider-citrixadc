package vpnglobal_vpneula_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalVpneulaBindingResourceModel describes the resource data model.
type VpnglobalVpneulaBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Eula                   types.String `tfsdk:"eula"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
}

func (r *VpnglobalVpneulaBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpneula_binding resource.",
			},
			"eula": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the EULA bound to vpnglobal",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// gotopriorityexpression is never echoed back by the NITRO GET
				// response for this binding, so it cannot be Computed (the value
				// would never resolve after apply -> "unknown value" error).
				// Keep it Optional-only; Read preserves the configured value
				// (Pattern 13 / Pattern 7).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
		},
	}
}

func vpnglobal_vpneula_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalVpneulaBindingResourceModel) vpn.Vpnglobalvpneulabinding {
	tflog.Debug(ctx, "In vpnglobal_vpneula_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_vpneula_binding := vpn.Vpnglobalvpneulabinding{}
	if !data.Eula.IsNull() && !data.Eula.IsUnknown() {
		vpnglobal_vpneula_binding.Eula = data.Eula.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_vpneula_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}

	return vpnglobal_vpneula_binding
}

func vpnglobal_vpneula_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpneulaBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpneulaBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpneula_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["eula"]; ok && val != nil {
		data.Eula = types.StringValue(val.(string))
	}
	// NOTE: gotopriorityexpression is a write-only / non-echoed input. The NITRO
	// GET response for this binding never returns it, so preserve the existing
	// plan/state value instead of nulling it out (Pattern 7).

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Eula.ValueString()))

	return data
}

// vpnglobal_vpneula_bindingSetAttrFromGetForDatasource faithfully copies the GET
// response into the model for the datasource flow (no prior plan/state to preserve).
func vpnglobal_vpneula_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalVpneulaBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpneulaBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpneula_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["eula"]; ok && val != nil {
		data.Eula = types.StringValue(val.(string))
	} else {
		data.Eula = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}

	// Set ID for the datasource (no Create runs in the datasource flow)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Eula.ValueString()))

	return data
}
