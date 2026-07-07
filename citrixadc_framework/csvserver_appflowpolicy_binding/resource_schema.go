package csvserver_appflowpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

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

// CsvserverAppflowpolicyBindingResourceModel describes the resource data model.
type CsvserverAppflowpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetlbvserver        types.String `tfsdk:"targetlbvserver"`
}

func (r *CsvserverAppflowpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_appflowpolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Invoke flag.",
			},
			"labelname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The invocation type.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policies bound to this vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Priority for the policy.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.\nExample: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1\nNote: Use this parameter only in case of Content Switching policy bind operations to a CS vserver",
			},
		},
	}
}

func csvserver_appflowpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *CsvserverAppflowpolicyBindingResourceModel) cs.Csvserverappflowpolicybinding {
	tflog.Debug(ctx, "In csvserver_appflowpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csvserver_appflowpolicy_binding := cs.Csvserverappflowpolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		csvserver_appflowpolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		csvserver_appflowpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		csvserver_appflowpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		csvserver_appflowpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		csvserver_appflowpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		csvserver_appflowpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		csvserver_appflowpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		csvserver_appflowpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
		csvserver_appflowpolicy_binding.Targetlbvserver = data.Targetlbvserver.ValueString()
	}

	return csvserver_appflowpolicy_binding
}

// csvserver_appflowpolicy_bindingSetAttrFromGet is the resource-side state setter.
// All schema attributes are RequiresReplace (no update endpoint), and several inputs
// (priority, gotopriorityexpression, invoke, labeltype, bindpoint) are server-defaulted
// or not faithfully echoed back. To avoid "inconsistent result after apply"/perpetual
// diffs it preserves the planned/state values rather than overwriting from the GET
// response. The ID is set once in Create and must not be recomputed here (Pattern 6/7).
func csvserver_appflowpolicy_bindingSetAttrFromGet(ctx context.Context, data *CsvserverAppflowpolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverAppflowpolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_appflowpolicy_bindingSetAttrFromGet Function")

	// The GET response for this binding echoes back name, policyname, priority,
	// gotopriorityexpression, invoke, labeltype, labelname. Copy those to resolve any
	// Computed unknowns. It does NOT echo bindpoint or targetlbvserver, so those
	// Optional+Computed values are PRESERVED from the plan/state (overwriting them with
	// null would cause an "inconsistent result after apply" / perpetual diff).
	// The ID is set once in Create and must NOT be recomputed here (Pattern 6/7).
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	// bindpoint / targetlbvserver: not echoed by GET — preserve plan/state value.
	// If still unknown (Computed and never set by user), resolve to null so the apply
	// result is consistent.
	if data.Bindpoint.IsUnknown() {
		data.Bindpoint = types.StringNull()
	}
	if data.Targetlbvserver.IsUnknown() {
		data.Targetlbvserver = types.StringNull()
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// csvserver_appflowpolicy_bindingSetAttrFromGetForDatasource faithfully copies the GET
// response into the model for the datasource (which has no prior plan/state to preserve)
// and sets the composite ID.
func csvserver_appflowpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *CsvserverAppflowpolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverAppflowpolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_appflowpolicy_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
