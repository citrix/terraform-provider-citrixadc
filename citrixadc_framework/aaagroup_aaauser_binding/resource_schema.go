package aaagroup_aaauser_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaagroupAaauserBindingResourceModel describes the resource data model.
type AaagroupAaauserBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Username               types.String `tfsdk:"username"`
}

func (r *AaagroupAaauserBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_aaauser_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The user name.",
			},
		},
	}
}

func aaagroup_aaauser_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaagroupAaauserBindingResourceModel) aaa.Aaagroupaaauserbinding {
	tflog.Debug(ctx, "In aaagroup_aaauser_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaagroup_aaauser_binding := aaa.Aaagroupaaauserbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaagroup_aaauser_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() {
		aaagroup_aaauser_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Username.IsNull() {
		aaagroup_aaauser_binding.Username = data.Username.ValueString()
	}

	return aaagroup_aaauser_binding
}

func aaagroup_aaauser_bindingSetAttrFromGet(ctx context.Context, data *AaagroupAaauserBindingResourceModel, getResponseData map[string]interface{}) *AaagroupAaauserBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_aaauser_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Groupname.ValueString(), data.Username.ValueString()))

	return data
}
