package vpnglobal_secureprivateaccessurl_binding

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

// VpnglobalSecureprivateaccessurlBindingResourceModel describes the resource data model.
type VpnglobalSecureprivateaccessurlBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Secureprivateaccessurl types.String `tfsdk:"secureprivateaccessurl"`
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_secureprivateaccessurl_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"secureprivateaccessurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Configured Secure Private Access URL",
			},
		},
	}
}

func vpnglobal_secureprivateaccessurl_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalSecureprivateaccessurlBindingResourceModel) vpn.Vpnglobalsecureprivateaccessurlbinding {
	tflog.Debug(ctx, "In vpnglobal_secureprivateaccessurl_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_secureprivateaccessurl_binding := vpn.Vpnglobalsecureprivateaccessurlbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_secureprivateaccessurl_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Secureprivateaccessurl.IsNull() && !data.Secureprivateaccessurl.IsUnknown() {
		vpnglobal_secureprivateaccessurl_binding.Secureprivateaccessurl = data.Secureprivateaccessurl.ValueString()
	}

	return vpnglobal_secureprivateaccessurl_binding
}

func vpnglobal_secureprivateaccessurl_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalSecureprivateaccessurlBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSecureprivateaccessurlBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_secureprivateaccessurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	// Pattern 7: the NITRO GET response does NOT echo back gotopriorityexpression
	// as a discrete field. Only overwrite it from the response when the key is
	// actually present; otherwise preserve the existing plan/state value so the
	// post-apply state matches the user config.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["secureprivateaccessurl"]; ok && val != nil {
		data.Secureprivateaccessurl = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessurl = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Secureprivateaccessurl.ValueString()))

	return data
}
