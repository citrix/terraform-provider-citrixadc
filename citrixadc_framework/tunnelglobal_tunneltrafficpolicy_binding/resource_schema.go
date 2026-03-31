package tunnelglobal_tunneltrafficpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/tunnel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TunnelglobalTunneltrafficpolicyBindingResourceModel describes the resource data model.
type TunnelglobalTunneltrafficpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tunnelglobal_tunneltrafficpolicy_binding resource.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
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
				Description: "Policy name.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Current state of the binding. If the binding is enabled, the policy is active.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind point to which the policy is bound.",
			},
		},
	}
}

func tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel) tunnel.Tunnelglobaltunneltrafficpolicybinding {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tunnelglobal_tunneltrafficpolicy_binding := tunnel.Tunnelglobaltunneltrafficpolicybinding{}
	if !data.Feature.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policyname.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.State.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.State = data.State.ValueString()
	}
	if !data.Type.IsNull() {
		tunnelglobal_tunneltrafficpolicy_binding.Type = data.Type.ValueString()
	}

	return tunnelglobal_tunneltrafficpolicy_binding
}

func tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel, getResponseData map[string]interface{}) *TunnelglobalTunneltrafficpolicyBindingResourceModel {
	tflog.Debug(ctx, "In tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
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
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
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
