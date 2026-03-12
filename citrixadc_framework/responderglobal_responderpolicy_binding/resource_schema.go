package responderglobal_responderpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ResponderglobalResponderpolicyBindingResourceModel describes the resource data model.
type ResponderglobalResponderpolicyBindingResourceModel struct {
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

func (r *ResponderglobalResponderpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderglobal_responderpolicy_binding resource.",
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
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Forward the request to the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the bind point whose policies you want to display. Available settings function as follows:\n* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT - Binds the policy to the default request queue.\n* OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue.\n* OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue..\n* SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue..\n* SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue.\n* RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue..\n* RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue.\n* MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue..\n* MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue.\n* MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue.\n* MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue.\n* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override response queue.\n* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default response queue.",
			},
		},
	}
}

func responderglobal_responderpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *ResponderglobalResponderpolicyBindingResourceModel) responder.Responderglobalresponderpolicybinding {
	tflog.Debug(ctx, "In responderglobal_responderpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderglobal_responderpolicy_binding := responder.Responderglobalresponderpolicybinding{}
	if !data.Globalbindtype.IsNull() {
		responderglobal_responderpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		responderglobal_responderpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		responderglobal_responderpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() {
		responderglobal_responderpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		responderglobal_responderpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		responderglobal_responderpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		responderglobal_responderpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		responderglobal_responderpolicy_binding.Type = data.Type.ValueString()
	}

	return responderglobal_responderpolicy_binding
}

func responderglobal_responderpolicy_bindingSetAttrFromGet(ctx context.Context, data *ResponderglobalResponderpolicyBindingResourceModel, getResponseData map[string]interface{}) *ResponderglobalResponderpolicyBindingResourceModel {
	tflog.Debug(ctx, "In responderglobal_responderpolicy_bindingSetAttrFromGet Function")

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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
