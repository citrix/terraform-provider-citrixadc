package sslglobal_sslpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

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

// SslglobalSslpolicyBindingResourceModel describes the resource data model.
type SslglobalSslpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *SslglobalSslpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslglobal_sslpolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				// Pattern 15: GET-only discriminator returned by the appliance.
				// Not accepted in the bind payload, so surface it as Computed-only.
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"invoke": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Invoke policies bound to a virtual server, service, or policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"labelname": schema.StringAttribute{
				// Pattern 8: only meaningful with invoke; CLI treats it as optional.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				// Pattern 8: only meaningful with invoke; CLI treats it as optional.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label to invoke. Specify virtual server for a policy label associated with a virtual server, or policy label for a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name for the SSL policy.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority of the policy binding.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Global bind point to which the policy is bound.",
			},
		},
	}
}

func sslglobal_sslpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslglobalSslpolicyBindingResourceModel) ssl.Sslglobalsslpolicybinding {
	tflog.Debug(ctx, "In sslglobal_sslpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslglobal_sslpolicy_binding := ssl.Sslglobalsslpolicybinding{}
	// Pattern 15: globalbindtype is a GET-only discriminator; excluded from the bind payload.
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		sslglobal_sslpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		sslglobal_sslpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		sslglobal_sslpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		sslglobal_sslpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		sslglobal_sslpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		sslglobal_sslpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		sslglobal_sslpolicy_binding.Type = data.Type.ValueString()
	}

	return sslglobal_sslpolicy_binding
}

func sslglobal_sslpolicy_bindingSetAttrFromGet(ctx context.Context, data *SslglobalSslpolicyBindingResourceModel, getResponseData map[string]interface{}) *SslglobalSslpolicyBindingResourceModel {
	tflog.Debug(ctx, "In sslglobal_sslpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	} else {
		data.Globalbindtype = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
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

	// Pattern 6: ID is set exactly once in Create; do not recompute it here.
	return data
}

// sslglobal_sslpolicy_bindingSetAttrFromGetForDatasource faithfully copies the GET
// response into the model and sets the composite ID, because the datasource has no
// Create step to establish it (Pattern 7 split).
func sslglobal_sslpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslglobalSslpolicyBindingResourceModel, getResponseData map[string]interface{}) *SslglobalSslpolicyBindingResourceModel {
	tflog.Debug(ctx, "In sslglobal_sslpolicy_bindingSetAttrFromGetForDatasource Function")

	sslglobal_sslpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

	// Set ID for the datasource: composite key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
