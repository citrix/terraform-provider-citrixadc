package aaauser_vpnsecureprivateaccessprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaauserVpnsecureprivateaccessprofileBindingResourceModel describes the resource data model.
type AaauserVpnsecureprivateaccessprofileBindingResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Gotopriorityexpression     types.String `tfsdk:"gotopriorityexpression"`
	Secureprivateaccessprofile types.String `tfsdk:"secureprivateaccessprofile"`
	Username                   types.String `tfsdk:"username"`
}

func (r *AaauserVpnsecureprivateaccessprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser_vpnsecureprivateaccessprofile_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"secureprivateaccessprofile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Secure Private Access Profile bound to the user.",
			},
			"username": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User account to which to bind the policy.",
			},
		},
	}
}

func aaauser_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaauserVpnsecureprivateaccessprofileBindingResourceModel) aaa.Aaauservpnsecureprivateaccessprofilebinding {
	tflog.Debug(ctx, "In aaauser_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaauser_vpnsecureprivateaccessprofile_binding := aaa.Aaauservpnsecureprivateaccessprofilebinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		aaauser_vpnsecureprivateaccessprofile_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Secureprivateaccessprofile.IsNull() && !data.Secureprivateaccessprofile.IsUnknown() {
		aaauser_vpnsecureprivateaccessprofile_binding.Secureprivateaccessprofile = data.Secureprivateaccessprofile.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		aaauser_vpnsecureprivateaccessprofile_binding.Username = data.Username.ValueString()
	}

	return aaauser_vpnsecureprivateaccessprofile_binding
}

func aaauser_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx context.Context, data *AaauserVpnsecureprivateaccessprofileBindingResourceModel, getResponseData map[string]interface{}) *AaauserVpnsecureprivateaccessprofileBindingResourceModel {
	tflog.Debug(ctx, "In aaauser_vpnsecureprivateaccessprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccessprofile"]; ok && val != nil {
		data.Secureprivateaccessprofile = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccessprofile = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
