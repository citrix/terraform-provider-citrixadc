package feoglobal_feopolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FeoglobalFeopolicyBindingResourceModel describes the resource data model.
type FeoglobalFeopolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *FeoglobalFeopolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the feoglobal_feopolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SYSTEM_GLOBAL"),
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the globally bound front end optimization policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "The priority assigned to the policy binding.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
		},
	}
}

func feoglobal_feopolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel) feo.Feoglobalfeopolicybinding {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	feoglobal_feopolicy_binding := feo.Feoglobalfeopolicybinding{}
	if !data.Globalbindtype.IsNull() {
		feoglobal_feopolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		feoglobal_feopolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policyname.IsNull() {
		feoglobal_feopolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		feoglobal_feopolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		feoglobal_feopolicy_binding.Type = data.Type.ValueString()
	}

	return feoglobal_feopolicy_binding
}

func feoglobal_feopolicy_bindingSetAttrFromGet(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel, getResponseData map[string]interface{}) *FeoglobalFeopolicyBindingResourceModel {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
