package aaauser_intranetip6_binding

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

// AaauserIntranetip6BindingResourceModel describes the resource data model.
type AaauserIntranetip6BindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetip6            types.String `tfsdk:"intranetip6"`
	Numaddr                types.Int64  `tfsdk:"numaddr"`
	Username               types.String `tfsdk:"username"`
}

func (r *AaauserIntranetip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser_intranetip6_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"intranetip6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Intranet IP6 bound to the user",
			},
			"numaddr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numbers of ipv6 address bound starting with intranetip6",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "User account to which to bind the policy.",
			},
		},
	}
}

func aaauser_intranetip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaauserIntranetip6BindingResourceModel) aaa.Aaauserintranetip6binding {
	tflog.Debug(ctx, "In aaauser_intranetip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaauser_intranetip6_binding := aaa.Aaauserintranetip6binding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaauser_intranetip6_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetip6.IsNull() {
		aaauser_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Numaddr.IsNull() {
		aaauser_intranetip6_binding.Numaddr = utils.IntPtr(int(data.Numaddr.ValueInt64()))
	}
	if !data.Username.IsNull() {
		aaauser_intranetip6_binding.Username = data.Username.ValueString()
	}

	return aaauser_intranetip6_binding
}

func aaauser_intranetip6_bindingSetAttrFromGet(ctx context.Context, data *AaauserIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *AaauserIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In aaauser_intranetip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
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
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Intranetip6.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Username.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
