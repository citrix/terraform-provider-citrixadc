package cmppolicylabel_cmppolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

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

// CmppolicylabelCmppolicyBindingResourceModel describes the resource data model.
type CmppolicylabelCmppolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invokelabelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *CmppolicylabelCmppolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmppolicylabel_cmppolicy_binding resource.",
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
				Description: "Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next higher priority number in the original label.",
			},
			"invokelabelname": schema.StringAttribute{
				// NITRO GET does not echo this field back, so it cannot be Computed
				// (would stay unknown after apply). Pure user input -> Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the label to invoke if the current policy evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the HTTP compression policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				// NITRO GET does not echo this field back, so it cannot be Computed
				// (would stay unknown after apply). Pure user input -> Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The compression policy name.",
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

func cmppolicylabel_cmppolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *CmppolicylabelCmppolicyBindingResourceModel) cmp.Cmppolicylabelcmppolicybinding {
	tflog.Debug(ctx, "In cmppolicylabel_cmppolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cmppolicylabel_cmppolicy_binding := cmp.Cmppolicylabelcmppolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		cmppolicylabel_cmppolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return cmppolicylabel_cmppolicy_binding
}

// cmppolicylabel_cmppolicy_bindingSetAttrFromGet updates the resource state from the
// GET response. The NITRO GET for this binding does NOT echo back labeltype or
// invokelabelname, so those (Optional+Computed) fields are preserved from the prior
// plan/state instead of being nulled (Pattern 7) to avoid "inconsistent result after
// apply". Fields that ARE returned by GET (gotopriorityexpression, invoke, priority,
// labelname, policyname) are set from the response.
func cmppolicylabel_cmppolicy_bindingSetAttrFromGet(ctx context.Context, data *CmppolicylabelCmppolicyBindingResourceModel, getResponseData map[string]interface{}) *CmppolicylabelCmppolicyBindingResourceModel {
	tflog.Debug(ctx, "In cmppolicylabel_cmppolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	// invokelabelname is not echoed back by the NITRO GET response; preserve the
	// existing plan/state value (Pattern 7) instead of nulling it.
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	// labeltype is not echoed back by the NITRO GET response; preserve the existing
	// plan/state value (Pattern 7) instead of nulling it.
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// cmppolicylabel_cmppolicy_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response for the datasource (which has no prior plan/state to
// preserve). Pattern 7 datasource split.
func cmppolicylabel_cmppolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *CmppolicylabelCmppolicyBindingResourceModel, getResponseData map[string]interface{}) *CmppolicylabelCmppolicyBindingResourceModel {
	tflog.Debug(ctx, "In cmppolicylabel_cmppolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (no Create runs for a datasource)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
