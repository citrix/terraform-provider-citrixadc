package aaagroup_vpnurl_binding

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

// AaagroupVpnurlBindingResourceModel describes the resource data model.
type AaagroupVpnurlBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Urlname                types.String `tfsdk:"urlname"`
}

func (r *AaagroupVpnurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_vpnurl_binding resource.",
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
			"urlname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet url",
			},
		},
	}
}

func aaagroup_vpnurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaagroupVpnurlBindingResourceModel) aaa.Aaagroupvpnurlbinding {
	tflog.Debug(ctx, "In aaagroup_vpnurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaagroup_vpnurl_binding := aaa.Aaagroupvpnurlbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaagroup_vpnurl_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() {
		aaagroup_vpnurl_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Urlname.IsNull() {
		aaagroup_vpnurl_binding.Urlname = data.Urlname.ValueString()
	}

	return aaagroup_vpnurl_binding
}

func aaagroup_vpnurl_bindingSetAttrFromGet(ctx context.Context, data *AaagroupVpnurlBindingResourceModel, getResponseData map[string]interface{}) *AaagroupVpnurlBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_vpnurl_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["urlname"]; ok && val != nil {
		data.Urlname = types.StringValue(val.(string))
	} else {
		data.Urlname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("urlname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Urlname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
