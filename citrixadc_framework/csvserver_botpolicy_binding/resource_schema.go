package csvserver_botpolicy_binding

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

// CsvserverBotpolicyBindingResourceModel describes the resource data model.
type CsvserverBotpolicyBindingResourceModel struct {
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

func (r *CsvserverBotpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_botpolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// Not echoed back by the NITRO GET response; keep as a pure
				// user input (no Computed) to avoid "unknown after apply".
				Optional: true,
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
				// Not echoed back by the NITRO GET response; pure user input.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				// Not echoed back by the NITRO GET response; pure user input.
				Optional: true,
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
				// Not echoed back by the NITRO GET response; pure user input.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.\nExample: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1\nNote: Use this parameter only in case of Content Switching policy bind operations to a CS vserver",
			},
		},
	}
}

func csvserver_botpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *CsvserverBotpolicyBindingResourceModel) cs.Csvserverbotpolicybinding {
	tflog.Debug(ctx, "In csvserver_botpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csvserver_botpolicy_binding := cs.Csvserverbotpolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		csvserver_botpolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		csvserver_botpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		csvserver_botpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		csvserver_botpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		csvserver_botpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		csvserver_botpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		csvserver_botpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		csvserver_botpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
		csvserver_botpolicy_binding.Targetlbvserver = data.Targetlbvserver.ValueString()
	}

	return csvserver_botpolicy_binding
}

// csvserver_botpolicy_bindingSetAttrFromGet is the resource-side state setter.
// The NITRO GET for a binding does not reliably echo back every configured input
// (e.g. priority/bindpoint/gotopriorityexpression/labeltype are server-defaulted or
// not echoed). To avoid "inconsistent result after apply" diffs, preserve the prior
// plan/state value when the GET response does not include a field; only adopt a value
// the server actually returns. The composite ID is set once in Create, not here.
func csvserver_botpolicy_bindingSetAttrFromGet(ctx context.Context, data *CsvserverBotpolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverBotpolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_botpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model. Preserve existing value when the field is
	// absent (or null) in the GET response.
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
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
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// csvserver_botpolicy_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (the datasource has no prior plan/state to preserve)
// and sets the composite ID itself, since the datasource never calls Create.
func csvserver_botpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *CsvserverBotpolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverBotpolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_botpolicy_bindingSetAttrFromGetForDatasource Function")

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
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}

	// Set ID for the datasource (no Create to set it)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
