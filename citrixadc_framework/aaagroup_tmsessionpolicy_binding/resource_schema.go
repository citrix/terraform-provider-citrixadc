package aaagroup_tmsessionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupTmsessionpolicyBindingResourceModel describes the resource data model.
type AaagroupTmsessionpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *AaagroupTmsessionpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup_tmsessionpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("REQUEST"),
				Description: "Bindpoint to which the policy is bound.",
			},
		},
	}
}

func aaagroup_tmsessionpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AaagroupTmsessionpolicyBindingResourceModel) aaa.Aaagrouptmsessionpolicybinding {
	tflog.Debug(ctx, "In aaagroup_tmsessionpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaagroup_tmsessionpolicy_binding := aaa.Aaagrouptmsessionpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		aaagroup_tmsessionpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupname.IsNull() {
		aaagroup_tmsessionpolicy_binding.Groupname = data.Groupname.ValueString()
	}
	if !data.Policy.IsNull() {
		aaagroup_tmsessionpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() {
		aaagroup_tmsessionpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		aaagroup_tmsessionpolicy_binding.Type = data.Type.ValueString()
	}

	return aaagroup_tmsessionpolicy_binding
}

func aaagroup_tmsessionpolicy_bindingSetAttrFromGet(ctx context.Context, data *AaagroupTmsessionpolicyBindingResourceModel, getResponseData map[string]interface{}) *AaagroupTmsessionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In aaagroup_tmsessionpolicy_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
