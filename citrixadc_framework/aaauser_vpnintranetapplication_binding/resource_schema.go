package aaauser_vpnintranetapplication_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaauserVpnintranetapplicationBindingResourceModel describes the resource data model.
type AaauserVpnintranetapplicationBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetapplication    types.String `tfsdk:"intranetapplication"`
	Username               types.String `tfsdk:"username"`
}

func (r *AaauserVpnintranetapplicationBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser_vpnintranetapplication_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"intranetapplication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the intranet VPN application to which the policy applies.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},
		},
	}
}

func aaauser_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaauserVpnintranetapplicationBindingResourceModel) aaa.Aaauservpnintranetapplicationbinding {
	tflog.Debug(ctx, "In aaauser_vpnintranetapplication_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaauser_vpnintranetapplication_binding := aaa.Aaauservpnintranetapplicationbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaauser_vpnintranetapplication_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetapplication.IsNull() {
		aaauser_vpnintranetapplication_binding.Intranetapplication = data.Intranetapplication.ValueString()
	}
	if !data.Username.IsNull() {
		aaauser_vpnintranetapplication_binding.Username = data.Username.ValueString()
	}

	return aaauser_vpnintranetapplication_binding
}

func aaauser_vpnintranetapplication_bindingSetAttrFromGet(ctx context.Context, data *AaauserVpnintranetapplicationBindingResourceModel, getResponseData map[string]interface{}) *AaauserVpnintranetapplicationBindingResourceModel {
	tflog.Debug(ctx, "In aaauser_vpnintranetapplication_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetapplication"]; ok && val != nil {
		data.Intranetapplication = types.StringValue(val.(string))
	} else {
		data.Intranetapplication = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetapplication:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetapplication.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
