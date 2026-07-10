package vpnglobal_vpnportaltheme_binding

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

// VpnglobalVpnportalthemeBindingResourceModel describes the resource data model.
type VpnglobalVpnportalthemeBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Portaltheme            types.String `tfsdk:"portaltheme"`
}

func (r *VpnglobalVpnportalthemeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnportaltheme_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// gotopriorityexpression is a write-only input: the NITRO GET response for
				// this binding never echoes it back. Dropping Computed (Pattern 8/13) avoids
				// "known after apply" / "inconsistent result" churn; the value is preserved
				// from plan/state in SetAttrFromGet (Pattern 7) instead of being nulled.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"portaltheme": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the portal theme bound to vpnglobal",
			},
		},
	}
}

func vpnglobal_vpnportaltheme_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalVpnportalthemeBindingResourceModel) vpn.Vpnglobalvpnportalthemebinding {
	tflog.Debug(ctx, "In vpnglobal_vpnportaltheme_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_vpnportaltheme_binding := vpn.Vpnglobalvpnportalthemebinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_vpnportaltheme_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Portaltheme.IsNull() && !data.Portaltheme.IsUnknown() {
		vpnglobal_vpnportaltheme_binding.Portaltheme = data.Portaltheme.ValueString()
	}

	return vpnglobal_vpnportaltheme_binding
}

// vpnglobal_vpnportaltheme_bindingSetAttrFromGet is the RESOURCE-side setter. It
// preserves user-supplied write-only inputs that the NITRO GET response never echoes
// (gotopriorityexpression) so Terraform does not see a spurious diff / inconsistent
// result after apply (Pattern 7). It does NOT recompute data.Id (Create sets it once).
func vpnglobal_vpnportaltheme_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnportalthemeBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnportalthemeBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnportaltheme_bindingSetAttrFromGet Function")

	// gotopriorityexpression is not echoed by the binding GET response; preserve the
	// existing plan/state value rather than nulling it (Pattern 7).
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["portaltheme"]; ok && val != nil {
		data.Portaltheme = types.StringValue(val.(string))
	}

	return data
}

// vpnglobal_vpnportaltheme_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior state to preserve, so it faithfully copies the
// GET response and sets data.Id itself (the datasource never calls Create) — Pattern 7.
func vpnglobal_vpnportaltheme_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalVpnportalthemeBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnportalthemeBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnportaltheme_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["portaltheme"]; ok && val != nil {
		data.Portaltheme = types.StringValue(val.(string))
	} else {
		data.Portaltheme = types.StringNull()
	}

	// Set ID for the datasource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Portaltheme.ValueString()))

	return data
}
