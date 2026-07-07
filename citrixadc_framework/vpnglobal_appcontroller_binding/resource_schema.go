package vpnglobal_appcontroller_binding

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

// VpnglobalAppcontrollerBindingResourceModel describes the resource data model.
type VpnglobalAppcontrollerBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Appcontroller          types.String `tfsdk:"appcontroller"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
}

func (r *VpnglobalAppcontrollerBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_appcontroller_binding resource.",
			},
			"appcontroller": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configured App Controller server.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Optional only (not Computed): this is a pure user input that the
				// NITRO GET never echoes back, so a Computed value could never be
				// resolved at apply time ("still indicated an unknown value").
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
		},
	}
}

func vpnglobal_appcontroller_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel) vpn.Vpnglobalappcontrollerbinding {
	tflog.Debug(ctx, "In vpnglobal_appcontroller_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_appcontroller_binding := vpn.Vpnglobalappcontrollerbinding{}
	if !data.Appcontroller.IsNull() && !data.Appcontroller.IsUnknown() {
		vpnglobal_appcontroller_binding.Appcontroller = data.Appcontroller.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_appcontroller_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}

	return vpnglobal_appcontroller_binding
}

// vpnglobal_appcontroller_bindingSetAttrFromGet is the RESOURCE state setter.
// It preserves user-configured / prior-state values for inputs that the NITRO GET
// response does not echo back (gotopriorityexpression is never returned by GET on
// this binding). It does NOT recompute data.Id — Create sets the ID exactly once.
func vpnglobal_appcontroller_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAppcontrollerBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_appcontroller_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appcontroller"]; ok && val != nil {
		data.Appcontroller = types.StringValue(val.(string))
	}
	// gotopriorityexpression is NOT echoed back by NITRO GET on this binding.
	// Preserve the existing plan/state value instead of nulling it (Pattern 7).

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Appcontroller.ValueString()))

	return data
}

// vpnglobal_appcontroller_bindingSetAttrFromGetForDatasource is the DATASOURCE setter.
// The datasource has no prior plan/state to preserve, so it faithfully copies every
// field from the GET response and sets the ID (the datasource has no Create).
func vpnglobal_appcontroller_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAppcontrollerBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_appcontroller_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["appcontroller"]; ok && val != nil {
		data.Appcontroller = types.StringValue(val.(string))
	} else {
		data.Appcontroller = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}

	// Set ID for the datasource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Appcontroller.ValueString()))

	return data
}
