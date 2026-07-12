package filterglobal_filterpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/filter"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FilterglobalFilterpolicyBindingResourceModel describes the resource data model.
type FilterglobalFilterpolicyBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Policyname types.String `tfsdk:"policyname"`
	Priority   types.Int64  `tfsdk:"priority"`
	State      types.String `tfsdk:"state"`
}

func (r *FilterglobalFilterpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the filterglobal_filterpolicy_binding resource.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the filter policy.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority of the policy.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "State of the binding.",
			},
		},
	}
}

func filterglobal_filterpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *FilterglobalFilterpolicyBindingResourceModel) filter.Filterglobalfilterpolicybinding {
	tflog.Debug(ctx, "In filterglobal_filterpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	filterglobal_filterpolicy_binding := filter.Filterglobalfilterpolicybinding{}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		filterglobal_filterpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		filterglobal_filterpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		filterglobal_filterpolicy_binding.State = data.State.ValueString()
	}

	return filterglobal_filterpolicy_binding
}

func filterglobal_filterpolicy_bindingSetAttrFromGet(ctx context.Context, data *FilterglobalFilterpolicyBindingResourceModel, getResponseData map[string]interface{}) *FilterglobalFilterpolicyBindingResourceModel {
	tflog.Debug(ctx, "In filterglobal_filterpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Single unique attribute - key:UrlEncode(value)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
