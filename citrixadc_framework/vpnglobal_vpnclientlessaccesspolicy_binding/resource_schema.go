package vpnglobal_vpnclientlessaccesspolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalVpnclientlessaccesspolicyBindingResourceModel describes the resource data model.
type VpnglobalVpnclientlessaccesspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Feature                types.String `tfsdk:"feature"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
	Type                   types.String `tfsdk:"type"`
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_vpnclientlessaccesspolicy_binding resource.",
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.",
			},
			"policyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.",
			},
			"secondary": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound",
			},
		},
	}
}

func vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel) vpn.Vpnglobalvpnclientlessaccesspolicybinding {
	tflog.Debug(ctx, "In vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_vpnclientlessaccesspolicy_binding := vpn.Vpnglobalvpnclientlessaccesspolicybinding{}
	if !data.Feature.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Policyname.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Secondary = data.Secondary.ValueBool()
	}
	if !data.Type.IsNull() {
		vpnglobal_vpnclientlessaccesspolicy_binding.Type = data.Type.ValueString()
	}

	return vpnglobal_vpnclientlessaccesspolicy_binding
}

func vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalVpnclientlessaccesspolicyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
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
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
