package botglobal_botpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotglobalBotpolicyBindingResourceModel describes the resource data model.
type BotglobalBotpolicyBindingResourceModel struct {
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

func (r *BotglobalBotpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botglobal_botpolicy_binding resource.",
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server.",
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
				Description: "Name of the bot policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the bind point whose policies you want to display. Available settings function as follows:\n* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT - Binds the policy to the default request queue.",
			},
		},
	}
}

func botglobal_botpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BotglobalBotpolicyBindingResourceModel) bot.Botglobalbotpolicybinding {
	tflog.Debug(ctx, "In botglobal_botpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botglobal_botpolicy_binding := bot.Botglobalbotpolicybinding{}
	if !data.Globalbindtype.IsNull() {
		botglobal_botpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		botglobal_botpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() {
		botglobal_botpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() {
		botglobal_botpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() {
		botglobal_botpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() {
		botglobal_botpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		botglobal_botpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() {
		botglobal_botpolicy_binding.Type = data.Type.ValueString()
	}

	return botglobal_botpolicy_binding
}

func botglobal_botpolicy_bindingSetAttrFromGet(ctx context.Context, data *BotglobalBotpolicyBindingResourceModel, getResponseData map[string]interface{}) *BotglobalBotpolicyBindingResourceModel {
	tflog.Debug(ctx, "In botglobal_botpolicy_bindingSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
