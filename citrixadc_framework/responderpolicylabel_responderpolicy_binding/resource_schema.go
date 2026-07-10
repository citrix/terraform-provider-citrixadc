package responderpolicylabel_responderpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

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

// ResponderpolicylabelResponderpolicyBindingResourceModel describes the resource data model.
type ResponderpolicylabelResponderpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invokelabelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderpolicylabel_responderpolicy_binding resource.",
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.",
			},
			"invokelabelname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the responder policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label to invoke. Available settings function as follows:\n* vserver - Invoke an unnamed policy label associated with a virtual server.\n* policylabel - Invoke a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the responder policy.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func responderpolicylabel_responderpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel) responder.Responderpolicylabelresponderpolicybinding {
	tflog.Debug(ctx, "In responderpolicylabel_responderpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	responderpolicylabel_responderpolicy_binding := responder.Responderpolicylabelresponderpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		responderpolicylabel_responderpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return responderpolicylabel_responderpolicy_binding
}

// responderpolicylabel_responderpolicy_bindingSetAttrFromGet is the RESOURCE-side
// state setter. Several inputs (gotopriorityexpression, invoke, invokelabelname,
// labeltype, priority) are server-overridden or not faithfully echoed by GET, so we
// PRESERVE the plan/state values to avoid "inconsistent result after apply" / perpetual
// diffs (Pattern 7/13). Identity keys (labelname, policyname) are adopted from the GET
// response. The ID is composed by Create; this setter does not recompute it.
func responderpolicylabel_responderpolicy_bindingSetAttrFromGet(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResponderpolicyBindingResourceModel {
	tflog.Debug(ctx, "In responderpolicylabel_responderpolicy_bindingSetAttrFromGet Function")

	// Identity keys: safe to adopt from the GET response.
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}

	// invokelabelname / labeltype are Optional+Computed server-assigned values (the
	// user typically does not set them). Adopt them from the GET response so the
	// Computed value is always KNOWN after apply (set null when absent rather than
	// leaving them unknown).
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
	}

	// gotopriorityexpression / invoke / priority: preserve the existing plan/state
	// value (these are user-driven inputs that the server may not echo or may
	// normalize — avoid "inconsistent result after apply"). Do NOT null them out.

	return data
}

// responderpolicylabel_responderpolicy_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter: it has no prior plan/state, so it faithfully copies every
// field from the GET response and sets the ID (Pattern 7 datasource split).
func responderpolicylabel_responderpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResponderpolicyBindingResourceModel {
	tflog.Debug(ctx, "In responderpolicylabel_responderpolicy_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
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

	// Set ID for the datasource: legacy 2-key composite (labelname, policyname)
	// matching resource_id_mapping.json.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
