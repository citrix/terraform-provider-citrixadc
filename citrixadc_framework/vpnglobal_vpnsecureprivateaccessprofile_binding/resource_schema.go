package vpnglobal_vpnsecureprivateaccessprofile_binding

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

// VpnglobalVpnsecureprivateaccessprofileBindingResourceModel describes the resource data model.
type VpnglobalVpnsecureprivateaccessprofileBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Gotopriorityexpression     types.String `tfsdk:"gotopriorityexpression"`
	Secureprivateaccessprofile types.String `tfsdk:"secureprivateaccessprofile"`
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnsecureprivateaccessprofile_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the Secure Private Access Profile bound to vpn global.",
			},
		},
	}
}

func vpnglobal_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalVpnsecureprivateaccessprofileBindingResourceModel) vpn.Vpnglobalvpnsecureprivateaccessprofilebinding {
	tflog.Debug(ctx, "In vpnglobal_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_vpnsecureprivateaccessprofile_binding := vpn.Vpnglobalvpnsecureprivateaccessprofilebinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_vpnsecureprivateaccessprofile_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Secureprivateaccessprofile.IsNull() && !data.Secureprivateaccessprofile.IsUnknown() {
		vpnglobal_vpnsecureprivateaccessprofile_binding.Secureprivateaccessprofile = data.Secureprivateaccessprofile.ValueString()
	}

	return vpnglobal_vpnsecureprivateaccessprofile_binding
}

func vpnglobal_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnsecureprivateaccessprofileBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnsecureprivateaccessprofileBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnsecureprivateaccessprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	// The NITRO GET response does NOT reliably echo back gotopriorityexpression
	// as a discrete field. Only overwrite it from the response when the key is
	// actually present; otherwise preserve the existing plan/state value so the
	// post-apply state matches the user config.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["secureprivateaccessprofile"]; ok && val != nil {
		data.Secureprivateaccessprofile = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))

	return data
}
