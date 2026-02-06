package vpnvserver_appfwpolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverAppfwpolicyBindingResourceModel describes the resource data model.
type VpnvserverAppfwpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Name                   types.String `tfsdk:"name"`
	Policy                 types.String `tfsdk:"policy"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *VpnvserverAppfwpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Binding object showing the appfwpolicy that can be bound to vpnvserver.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Terraform ID. Format: <name>,<policy>",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "The name of the policy, if any, bound to the VPN virtual server.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.",
			},
		},
	}
}

func vpnvserverAppfwpolicyBindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverAppfwpolicyBindingResourceModel) vpn.Vpnvserverappfwpolicybinding {
	tflog.Debug(ctx, "In vpnvserverAppfwpolicyBindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserverAppfwpolicyBinding := vpn.Vpnvserverappfwpolicybinding{}

	if !data.Name.IsNull() {
		vpnvserverAppfwpolicyBinding.Name = data.Name.ValueString()
	}
	if !data.Policy.IsNull() {
		vpnvserverAppfwpolicyBinding.Policy = data.Policy.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		vpnvserverAppfwpolicyBinding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() {
		vpnvserverAppfwpolicyBinding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Priority.IsNull() {
		vpnvserverAppfwpolicyBinding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() {
		vpnvserverAppfwpolicyBinding.Secondary = data.Secondary.ValueBool()
	}

	return vpnvserverAppfwpolicyBinding
}

func vpnvserverAppfwpolicyBindingSetAttrFromGet(ctx context.Context, data *VpnvserverAppfwpolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverAppfwpolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserverAppfwpolicyBindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		if boolVal, ok := val.(bool); ok {
			data.Groupextraction = types.BoolValue(boolVal)
		}
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		if boolVal, ok := val.(bool); ok {
			data.Secondary = types.BoolValue(boolVal)
		}
	} else {
		data.Secondary = types.BoolNull()
	}

	return data
}
