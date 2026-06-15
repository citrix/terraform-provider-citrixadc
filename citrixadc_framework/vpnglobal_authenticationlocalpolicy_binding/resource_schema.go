package vpnglobal_authenticationlocalpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

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

// VpnglobalAuthenticationlocalpolicyBindingResourceModel describes the resource data model.
type VpnglobalAuthenticationlocalpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnglobalAuthenticationlocalpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_authenticationlocalpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// Optional-only (no Computed): user-driven input that the NITRO GET response
				// does not echo back. A Computed flag would leave it "known after apply" and
				// cause inconsistent-result-after-apply errors (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				// Optional-only (no Computed): not echoed by the NITRO GET response (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				// Optional-only (no Computed): user-driven RequiresReplace input (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				// Optional-only (no Computed): user-driven RequiresReplace input (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.",
			},
		},
	}
}

func vpnglobal_authenticationlocalpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalAuthenticationlocalpolicyBindingResourceModel) vpn.Vpnglobalauthenticationlocalpolicybinding {
	tflog.Debug(ctx, "In vpnglobal_authenticationlocalpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_authenticationlocalpolicy_binding := vpn.Vpnglobalauthenticationlocalpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_authenticationlocalpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		vpnglobal_authenticationlocalpolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		vpnglobal_authenticationlocalpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		vpnglobal_authenticationlocalpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		vpnglobal_authenticationlocalpolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return vpnglobal_authenticationlocalpolicy_binding
}

// vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGet is the RESOURCE-side state
// setter. The NITRO GET response for this binding does NOT echo back every input
// (notably gotopriorityexpression and groupextraction, and secondary/priority are only
// returned when set), so this setter PRESERVES the existing plan/state value when a field
// is absent from the response instead of nulling it. Nulling non-echoed user inputs would
// produce inconsistent-result-after-apply errors (Pattern 7, server-overrides/non-echoed
// inputs variant). All attributes are RequiresReplace, so preserving state is correct.
func vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalAuthenticationlocalpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationlocalpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGet Function")

	// All non-key attributes are Optional-only (no Computed) and RequiresReplace. For such
	// attributes Terraform requires the post-apply value to EXACTLY equal the configured
	// value, so the resource-side setter must NOT overwrite them from the GET response —
	// the NITRO server echoes fields it was never asked to set (e.g. secondary:false even
	// when the user left it null), which would trigger "inconsistent result after apply".
	// We therefore preserve the planned/state values verbatim for every non-key attribute
	// and only refresh the key (policyname) + ID from the response.
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}

// vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter (Pattern 7 datasource split). A datasource has no prior
// plan/state to preserve, so it faithfully copies every field from the GET response
// (nulling absent fields) and sets data.Id, since the datasource has no Create.
func vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalAuthenticationlocalpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalAuthenticationlocalpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_authenticationlocalpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Datasource has no Create — set ID here (single unique attr → plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
