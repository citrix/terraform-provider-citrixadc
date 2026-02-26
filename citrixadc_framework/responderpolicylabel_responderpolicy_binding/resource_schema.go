package responderpolicylabel_responderpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ResponderpolicylabelResponderpolicyBindingResourceModel describes the resource data model.
type ResponderpolicylabelResponderpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
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
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke. Available settings function as follows:\n* vserver - Invoke an unnamed policy label associated with a virtual server.\n* policylabel - Invoke a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func responderpolicylabel_responderpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel) responder.Responderpolicylabelresponderpolicybinding {
	tflog.Debug(ctx, "In responderpolicylabel_responderpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderpolicylabel_responderpolicy_binding := responder.Responderpolicylabelresponderpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		responderpolicylabel_responderpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		responderpolicylabel_responderpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() {
		responderpolicylabel_responderpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() {
		responderpolicylabel_responderpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		responderpolicylabel_responderpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		responderpolicylabel_responderpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		responderpolicylabel_responderpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return responderpolicylabel_responderpolicy_binding
}

func responderpolicylabel_responderpolicy_bindingSetAttrFromGet(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResponderpolicyBindingResourceModel {
	tflog.Debug(ctx, "In responderpolicylabel_responderpolicy_bindingSetAttrFromGet Function")

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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
