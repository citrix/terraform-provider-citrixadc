package gslbvserver_lbpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbvserverLbpolicyBindingResourceModel describes the resource data model.
type GslbvserverLbpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Name                   types.String `tfsdk:"name"`
	Order                  types.Int64  `tfsdk:"order"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *GslbvserverLbpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbvserver_lbpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.\n	o	If gotoPriorityExpression is not present or if it is equal to END then the policy bank evaluation ends here\n	o	Else if the gotoPriorityExpression is equal to NEXT then the next policy in the priority order is evaluated.\n	o	Else gotoPriorityExpression is evaluated. The result of gotoPriorityExpression (which has to be a number) is processed as follows:\n		-	An UNDEF event is triggered if\n			.	gotoPriorityExpression cannot be evaluated\n			.	gotoPriorityExpression evaluates to number which is smaller than the maximum priority in the policy bank but is not same as any policy's priority\n			.	gotoPriorityExpression evaluates to a priority that is smaller than the current policy's priority\n		-	If the gotoPriorityExpression evaluates to the priority of the current policy then the next policy in the priority order is evaluated.\n		-	If the gotoPriorityExpression evaluates to the priority of a policy further ahead in the list then that policy will be evaluated next.\n		This field is applicable only to rewrite and responder policies.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy bound to the GSLB vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The bindpoint to which the policy is bound",
			},
		},
	}
}

func gslbvserver_lbpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbvserverLbpolicyBindingResourceModel) gslb.Gslbvserverlbpolicybinding {
	tflog.Debug(ctx, "In gslbvserver_lbpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbvserver_lbpolicy_binding := gslb.Gslbvserverlbpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		gslbvserver_lbpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Name.IsNull() {
		gslbvserver_lbpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() {
		gslbvserver_lbpolicy_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Policyname.IsNull() {
		gslbvserver_lbpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		gslbvserver_lbpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		gslbvserver_lbpolicy_binding.Type = data.Type.ValueString()
	}

	return gslbvserver_lbpolicy_binding
}

func gslbvserver_lbpolicy_bindingSetAttrFromGet(ctx context.Context, data *GslbvserverLbpolicyBindingResourceModel, getResponseData map[string]interface{}) *GslbvserverLbpolicyBindingResourceModel {
	tflog.Debug(ctx, "In gslbvserver_lbpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
