package aaagroup_intranetip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupIntranetip6BindingResourceModel describes the resource data model.
type AaagroupIntranetip6BindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Intranetip6            types.String `tfsdk:"intranetip6"`
	Numaddr                types.Int64  `tfsdk:"numaddr"`
}

func (r *AaagroupIntranetip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_intranetip6_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the group that you are binding.",
			},
			"intranetip6": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The Intranet IP6(s) bound to the group",
			},
			"numaddr": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Numbers of ipv6 address bound starting with intranetip6",
			},
		},
	}
}

func aaagroup_intranetip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel) aaa.Aaagroupintranetip6binding {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaagroup_intranetip6_binding := aaa.Aaagroupintranetip6binding{}
	// Pattern 15: gotopriorityexpression is exclusive to the -policy branch of
	// `bind aaa group` and is NOT accepted on the intranetIP6 branch. Never send it.
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaagroup_intranetip6_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Intranetip6.IsNull() && !data.Intranetip6.IsUnknown() {
		aaagroup_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Numaddr.IsNull() && !data.Numaddr.IsUnknown() {
		aaagroup_intranetip6_binding.Numaddr = utils.IntPtr(int(data.Numaddr.ValueInt64()))
	}

	return aaagroup_intranetip6_binding
}

// aaagroup_intranetip6_bindingSetAttrFromGet is the resource-side setter.
// It preserves user-configured values (gotopriorityexpression is never echoed
// in a meaningful way and is not part of this branch) and only refreshes the
// identity attributes from the GET response. It does NOT (re)compute the ID -
// the resource sets the ID exactly once in Create.
func aaagroup_intranetip6_bindingSetAttrFromGet(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *AaagroupIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingSetAttrFromGet Function")

	// gotopriorityexpression is not sent/echoed on the intranetIP6 branch -
	// preserve the existing plan/state value, do not overwrite it.
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	}

	return data
}

// aaagroup_intranetip6_bindingSetAttrFromGetForDatasource faithfully copies
// every field from the GET response and composes the ID, because the datasource
// never calls Create and has no prior state to preserve.
func aaagroup_intranetip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AaagroupIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *AaagroupIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_intranetip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	} else {
		data.Intranetip6 = types.StringNull()
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	} else {
		data.Numaddr = types.Int64Null()
	}

	// Compose the composite ID: groupname,intranetip6,numaddr.
	// IPv6 colons are percent-encoded by UrlEncode so they never collide with
	// the key:value / comma delimiters used by ParseIdString.
	data.Id = types.StringValue(aaagroup_intranetip6_bindingComposeId(
		data.Groupname.ValueString(),
		data.Intranetip6.ValueString(),
		data.Numaddr.ValueInt64(),
	))

	return data
}

// aaagroup_intranetip6_bindingComposeId builds the colon-safe composite ID.
// Each value is UrlEncoded so the IPv6 address's literal ':' characters become
// '%3A' and cannot be mistaken for the key:value separator that ParseIdString
// keys on (it finds the FIRST ':' in each comma-separated segment).
func aaagroup_intranetip6_bindingComposeId(groupname string, intranetip6 string, numaddr int64) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(groupname)))
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.UrlEncode(intranetip6)))
	idParts = append(idParts, fmt.Sprintf("numaddr:%s", utils.UrlEncode(fmt.Sprintf("%d", numaddr))))
	return strings.Join(idParts, ",")
}
