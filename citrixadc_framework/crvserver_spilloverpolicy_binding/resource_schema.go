package crvserver_spilloverpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CrvserverSpilloverpolicyBindingResourceModel describes the resource data model.
type CrvserverSpilloverpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetvserver          types.String `tfsdk:"targetvserver"`
}

func (r *CrvserverSpilloverpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crvserver_spilloverpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke a policy label if this policy's rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to be invoked.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of label to be invoked.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policies bound to this vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority for the policy.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.",
			},
		},
	}
}

func crvserver_spilloverpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *CrvserverSpilloverpolicyBindingResourceModel) cr.Crvserverspilloverpolicybinding {
	tflog.Debug(ctx, "In crvserver_spilloverpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	crvserver_spilloverpolicy_binding := cr.Crvserverspilloverpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() {
		crvserver_spilloverpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		crvserver_spilloverpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() {
		crvserver_spilloverpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		crvserver_spilloverpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Name.IsNull() {
		crvserver_spilloverpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policyname.IsNull() {
		crvserver_spilloverpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		crvserver_spilloverpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Targetvserver.IsNull() {
		crvserver_spilloverpolicy_binding.Targetvserver = data.Targetvserver.ValueString()
	}

	return crvserver_spilloverpolicy_binding
}

func crvserver_spilloverpolicy_bindingSetAttrFromGet(ctx context.Context, data *CrvserverSpilloverpolicyBindingResourceModel, getResponseData map[string]interface{}) *CrvserverSpilloverpolicyBindingResourceModel {
	tflog.Debug(ctx, "In crvserver_spilloverpolicy_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["targetvserver"]; ok && val != nil {
		data.Targetvserver = types.StringValue(val.(string))
	} else {
		data.Targetvserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
