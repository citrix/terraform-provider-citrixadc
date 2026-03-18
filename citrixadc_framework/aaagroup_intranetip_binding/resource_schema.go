package aaagroup_intranetip_binding

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

// AaagroupIntranetipBindingResourceModel describes the resource data model.
type AaagroupIntranetipBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Intranetip             types.String `tfsdk:"intranetip"`
	Netmask                types.String `tfsdk:"netmask"`
}

func (r *AaagroupIntranetipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_intranetip_binding resource.",
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
			"intranetip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Intranet IP(s) bound to the group",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask for the Intranet IP",
			},
		},
	}
}

func aaagroup_intranetip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaagroupIntranetipBindingResourceModel) aaa.Aaagroupintranetipbinding {
	tflog.Debug(ctx, "In aaagroup_intranetip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaagroup_intranetip_binding := aaa.Aaagroupintranetipbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaagroup_intranetip_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() {
		aaagroup_intranetip_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Intranetip.IsNull() {
		aaagroup_intranetip_binding.Intranetip = data.Intranetip.ValueString()
	}
	if !data.Netmask.IsNull() {
		aaagroup_intranetip_binding.Netmask = data.Netmask.ValueString()
	}

	return aaagroup_intranetip_binding
}

func aaagroup_intranetip_bindingSetAttrFromGet(ctx context.Context, data *AaagroupIntranetipBindingResourceModel, getResponseData map[string]interface{}) *AaagroupIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_intranetip_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	} else {
		data.Intranetip = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
